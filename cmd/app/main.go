package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"AATool_Tracker/internal/minecraft"
	"AATool_Tracker/internal/models"
	"AATool_Tracker/internal/tracker"
)

func main() {
	fmt.Println("=== AATool (Go Edition) ===")

	savesPath := minecraft.GetSavesPath()
	worldPath, err := minecraft.GetLatestWorld(savesPath)
	if err != nil {
		exitWithError(err)
	}

	fmt.Printf("Выбран мир: %s\n", filepath.Base(worldPath))
	fmt.Println("Ожидание файла достижений...")

	var advFile string
	for {
		advFile, err = minecraft.FindAdvancementFile(worldPath)
		if err == nil && advFile != "" {
			break
		}
		time.Sleep(2 * time.Second)
	}

	fmt.Printf("Файл найден! Отслеживание запущено.\n\n")

	// Создаем канал для связи логики и интерфейса
	eventsChan := make(chan models.AdvancementEvent)

	// Запускаем трекер в отдельном потоке (goroutine)
	go tracker.Watch(advFile, eventsChan)

	// В главном потоке слушаем канал
	// Этот цикл будет ждать, пока в канал не прилетят данные
	for event := range eventsChan {
		fmt.Printf("[%s] ПОЛУЧЕНО: %s\n", event.Time.Format("02.01.2006 15:04:05"), event.AchievementDisplayName)
	}
}

func exitWithError(err error) {
	fmt.Printf("ОШИБКА: %v\n", err)
	fmt.Println("Нажмите Enter для выхода...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	os.Exit(1)
}
