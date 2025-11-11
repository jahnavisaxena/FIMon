package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func CreateBaseline(dir string, baselinePath string) map[string]string {
	baseline := make(map[string]string)

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		baseline[path] = GetFileHash(path)
		return nil
	})

	data, _ := json.MarshalIndent(baseline, "", "  ")
	os.WriteFile(baselinePath, data, 0644)
	fmt.Println("[+] Baseline created at:", baselinePath)
	return baseline
}

func LoadBaseline(baselinePath string) map[string]string {
	baseline := make(map[string]string)
	file, err := os.Open(baselinePath)
	if err != nil {
		return baseline
	}
	defer file.Close()

	json.NewDecoder(file).Decode(&baseline)
	return baseline
}

// SaveBaseline writes the updated baseline data to a JSON file
func SaveBaseline(baseline map[string]string, baselinePath string) {
	data, err := json.MarshalIndent(baseline, "", "  ")
	if err != nil {
		fmt.Println("Error encoding baseline:", err)
		return
	}

	err = os.WriteFile(baselinePath, data, 0644)
	if err != nil {
		fmt.Println("Error writing baseline file:", err)
		return
	}

	fmt.Println("[✔] Baseline updated:", baselinePath)
}

