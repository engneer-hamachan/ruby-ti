package builtin

import (
	"embed"
	"log"
)

//go:embed builtin_config
var configFiles embed.FS

func init() {
	if err := loadBuiltinFromJSON(configFiles, "builtin_config"); err != nil {
		log.Fatalf("Failed to load builtin configurations: %v", err)
	}
}
