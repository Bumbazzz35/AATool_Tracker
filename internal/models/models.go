package models

import "time"

// AdvancementEvent - событие получения достижения
type AdvancementEvent struct {
	ID                     string    // Системный ID (например, "minecraft:story/mine_stone")
	AchievementDisplayName string    // Красивое имя ("Каменный век")
	Time                   time.Time // Время получения
	IconPath               string    // путь к иконке для GUI
}
