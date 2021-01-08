package main

import (
	"./BotLogger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"os"
)

type Bot struct {
	commonLogger *BotLogger.BotLogger
	errLogger    *BotLogger.BotLogger
	bot          *tgbotapi.BotAPI
}

func (b *Bot) Setup(token string) {
	b.commonLogger = new(BotLogger.BotLogger)
	b.errLogger = new(BotLogger.BotLogger)
	b.bot = new(tgbotapi.BotAPI)

	var err error
	b.commonLogger.Setup("")
	b.errLogger.Setup("error.log")

	b.bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		b.errLogger.Outputln("Failed to Create a Bot API Instance.", err)
	}
	b.commonLogger.Outputf("Authorized on account %s", b.bot.Self.UserName)
}

func main() {
	TestBot := new(Bot)
	TestBot.Setup(os.Getenv("TeleBotToken"))

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := TestBot.bot.GetUpdatesChan(u)
	if err != nil {
		TestBot.errLogger.Outputln("Failed to Get Updates.", err)
		os.Exit(-1)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}
		TestBot.commonLogger.Outputf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	}
}
