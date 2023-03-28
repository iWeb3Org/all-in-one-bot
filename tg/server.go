package tg

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/uerax/goconf"
)

func Server() {

	// Create a new bot instance
	token, err := goconf.VarString("telegram", "token")
	if err != nil {
		panic(err)
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	// Set bot options
	bot.Debug = true
	
	 
	// Create a new update channel
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	api.New()

	// Start listening for updates
	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil { // ignore non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
            continue
        }

		switch update.Message.Command() {
		case "addCryptoMoniter":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%v", api.cryptoApi.Price(update.Message.CommandArguments())))
			_, err := bot.Send(msg)
			if err != nil {
				log.Println(err)
			}
		}
		
	}
}
