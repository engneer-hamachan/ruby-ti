package loader

import (
	"encoding/json"
	"fmt"
	"os"
)

type TiLoader struct {
	Preload []string `json:"preload"`
}

func GetPreloadFiles() []string {
	configPath := ".ti-loader.json"

	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}
		}
		fmt.Fprintf(os.Stderr, "Error reading %s: %v\n", configPath, err)
		os.Exit(1)
	}

	var loader TiLoader
	if err := json.Unmarshal(data, &loader); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing %s: %v\n", configPath, err)
		os.Exit(1)
	}

	return loader.Preload
}
