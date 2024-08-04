package components

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

const (
	ProfileDataPath = "data/profile_info.json"
	IconsDataPath   = "data/icons_info.json"
)

func ReadJsonFile(path string) data {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("could not get current working directory %v", err)
	}
	b, err := os.ReadFile(strings.Join([]string{wd, path}, `/`))
	if err != nil {
		log.Fatalf("could not open json file %s, %v", path, err)
	}
	var d data
	err = json.Unmarshal(b, &d)
	if err != nil {
		log.Fatalf("could not unmarshal json data %s", err)
	}
	return d
}
