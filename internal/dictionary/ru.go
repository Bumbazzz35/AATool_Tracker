package dictionary

import "strings"

var ruNames = map[string]string{
	"minecraft:story/mine_stone":       "Каменный век",
	"minecraft:story/upgrade_tools":    "Обновка!",
	"minecraft:story/smelt_iron":       "Куй железо",
	"minecraft:story/obtain_armor":     "Дресс-код",
	"minecraft:story/lava_bucket":      "Горячая штучка",
	"minecraft:story/iron_tools":       "И кирка без дела не ржавеет",
	"minecraft:story/mine_diamond":     "Алмазы!",
	"minecraft:story/enter_the_nether": "Огненные недра",
}

// Translate преобразует ID в красивое название
func Translate(id string) string {
	if name, ok := ruNames[id]; ok {
		return name
	}

	// Если перевода нет, делаем из ID читаемый текст
	parts := strings.Split(id, "/")
	if len(parts) > 0 {
		return strings.ReplaceAll(parts[len(parts)-1], "_", " ")
	}
	return id
}
