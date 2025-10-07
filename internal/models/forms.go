package models

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var formQuestions []string

func LoadFormQuestions(file string) []string {

	var path = os.Getenv("QUESTIONS_PATH") + file + ".yaml"

	yamlFile, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error while reading file: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, &formQuestions)
	if err != nil {
		log.Fatalf("Error while reading file: %v", err)
	}

	return formQuestions
}
