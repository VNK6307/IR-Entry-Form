package main

import (
	"fmt"
	"log"
	"racer/form/config"
	"racer/form/internal/models"
)

func main() {

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Could not load configuration", err)
	}

	fmt.Println("Bot's token: ", cfg.BotToken)
	fmt.Println("Bot's address:", cfg.BotURL)

	questions := models.LoadFormQuestions()

	for _, question := range questions {
		fmt.Println(question)
	}

}
