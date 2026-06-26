package minecraft

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// GetSavesPath возвращает путь к папке saves
func GetSavesPath() string {
	// Если используешь кастомный лаунчер, можешь жестко вписать путь сюда
	// return "C:\\Users\\ТВОЕ_ИМЯ\\Путь_К_Лаунчеру\\saves"

	appdata := os.Getenv("APPDATA")
	if appdata == "" {
		return ".minecraft/saves"
	}
	return filepath.Join(appdata, ".minecraft", "saves")
}

// GetLatestWorld находит последний измененный мир
func GetLatestWorld(savesPath string) (string, error) {
	entries, err := os.ReadDir(savesPath)
	if err != nil {
		return "", err
	}

	var latestWorld string
	var latestModTime time.Time

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		info, err := entry.Info()
		if err != nil {
			continue
		}
		if info.ModTime().After(latestModTime) {
			latestModTime = info.ModTime()
			latestWorld = filepath.Join(savesPath, entry.Name())
		}
	}

	if latestWorld == "" {
		return "", fmt.Errorf("миры не найдены в папке %s", savesPath)
	}
	return latestWorld, nil
}

// FindAdvancementFile ищет .json файл игрока
func FindAdvancementFile(worldPath string) (string, error) {
	paths := []string{
		filepath.Join(worldPath, "advancements"),
		filepath.Join(worldPath, "players", "advancements"),
	}

	for _, p := range paths {
		entries, err := os.ReadDir(p)
		if err != nil {
			continue
		}
		for _, entry := range entries {
			if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".json") {
				return filepath.Join(p, entry.Name()), nil
			}
		}
	}
	return "", fmt.Errorf("файл достижений пока не создан (зайдите в мир)")
}
