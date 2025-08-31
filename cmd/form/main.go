package main

import (
	"fmt"
	"log"
	"racer/form/config"
)

func main() {

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Could not load configuration", err)
	}

	fmt.Println("Bot's token: ", cfg.BotToken) // ToDo Delete after bot has done
	fmt.Println("Bot's address:", cfg.BotURL)  // ToDo Delete after bot has done

}
