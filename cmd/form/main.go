package main

import (
	"fmt"
	"log"
	"racer/form/internal/config"
	"racer/form/internal/models"
)

func main() {

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Could not load configuration", err)
	}

	fmt.Println(cfg.BotToken)

	questions := models.LoadFormQuestions()

	for _, question := range questions {
		fmt.Println(question)
	}

}
