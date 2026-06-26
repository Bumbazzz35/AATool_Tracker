package tracker

import (
	"encoding/json"
	"os"
	"strings"
	"time"

	"AATool_Tracker/internal/dictionary"
	"AATool_Tracker/internal/models"
)

// Watch запускает бесконечный цикл проверки файла
func Watch(filePath string, eventsChan chan<- models.AdvancementEvent) {
	var lastModTime time.Time
	knownCompleted := make(map[string]bool)
	firstLoad := true

	for {
		info, err := os.Stat(filePath)
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}

		if info.ModTime().After(lastModTime) {
			lastModTime = info.ModTime()
			processFile(filePath, knownCompleted, firstLoad, eventsChan)
			firstLoad = false
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func processFile(filePath string, known map[string]bool, firstLoad bool, outChan chan<- models.AdvancementEvent) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return
	}

	var root map[string]interface{}
	if err := json.Unmarshal(data, &root); err != nil {
		return
	}

	for key, val := range root {
		if key == "DataVersion" || strings.Contains(key, "recipes/") || strings.HasSuffix(key, "root") {
			continue
		}

		advMap, ok := val.(map[string]interface{})
		if !ok {
			continue
		}

		done, exists := advMap["done"]
		if !exists {
			continue
		}

		isDone, ok := done.(bool)
		if ok && isDone && !known[key] {
			known[key] = true

			if !firstLoad {
				// Отправляем событие в канал
				outChan <- models.AdvancementEvent{
					ID:                     key,
					AchievementDisplayName: dictionary.Translate(key),
					Time:                   time.Now(),
				}
			}
		}
	}
}
