package dictionary

import (
	_ "embed"
	"encoding/json"
	"log"
	"strings"
)

//go:embed lang_ru.json
var langFile []byte

var achievementsNames map[string]string

func init() {
	if err := json.Unmarshal(langFile, &achievementsNames); err != nil {
		log.Printf("failed to decode embedded language file: %v", err)
	}
}

// Translate преобразует ID в красивое название
func Translate(id string) string {
	if achievementsNames != nil {
		if name, ok := achievementsNames[id]; ok {
			return name
		}
	}
	return formatID(id)
}

func formatID(id string) string {
	parts := strings.Split(id, "/")
	name := parts[len(parts)-1]
	return strings.ReplaceAll(name, "_", " ")
}
