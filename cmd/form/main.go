package main

import (
	"fmt"
	"log"
	"racer/form/config"
	"racer/form/internal/services"
)

type Bot struct {
	telegram services.TelegramService
	//handler  *handlers.Handler
}

func main() {

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Could not load configuration", err)
	}

	fmt.Println("Bot's token: ", cfg.BotToken) // ToDo Delete after bot has done
	fmt.Println("Bot's address:", cfg.BotURL)  // ToDo Delete after bot has done

	tgService := services.NewTelegramService(cfg)

	bot := &Bot{telegram: tgService}
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
			//b.handler.HandleUpdate(u)
			offset = update.UpdateID + 1
		}

	}
}
