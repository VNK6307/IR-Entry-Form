package models

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

const formSize int = 12 // ToDo Переделать в получение из файла env

var formQuestions [formSize]string

func LoadFormQuestions() [formSize]string {

	var path = os.Getenv("QUESTIONS_PATH")

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
