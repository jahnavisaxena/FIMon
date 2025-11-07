
package main

import (
	"encoding/json"
	"os"
)

func LoadBaseline(file string) (map[string]string, error) {
	baseline := make(map[string]string)
	f, err := os.Open(file)
	if err != nil {
		return baseline, err
	}
	defer f.Close()
	err = json.NewDecoder(f).Decode(&baseline)
	return baseline, err
}

func SaveBaseline(file string, baseline map[string]string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(baseline)
}
