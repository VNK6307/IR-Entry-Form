package main

import (
	"fmt"
	"log"
	"racer/form/config"
	"racer/form/internal/handlers"
	"racer/form/internal/repository"
	"racer/form/internal/services"
)

type Bot struct {
	telegram services.TelegramService
	handler  *handlers.Handler
}

func main() {

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Could not load configuration", err)
	}

	teamRepo := repository.NewTeamRepository()

	tgService := services.NewTelegramService(cfg)
	handler := handlers.NewHandler(tgService, teamRepo)

	bot := &Bot{tgService, handler}
	bot.Start()

}

func (bot *Bot) Start() {
	offset := 0
	for {
		updates, err := bot.telegram.GetUpdates(offset)
		if err != nil {
			log.Println("getUpdates error:", err)
			continue
		}
		for _, update := range updates {
			bot.handler.HandleUpdate(update)
			offset = update.UpdateID + 1
		}

	}
}
