package TeleBot

import (
	"TeleBot/BotLogger"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"os"
)

type Bot struct {
	commonLogger *BotLogger.BotLogger
	errLogger    *BotLogger.BotLogger
	bot          *tgbotapi.BotAPI
}

func (b *Bot) Setup(token string) {
	var err error
	b.commonLogger.Setup("")
	b.errLogger.Setup("error.log")

	b.bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		b.errLogger.Outputln("Failed to Create a Bot API Instance. ", err)
	}
	b.commonLogger.Outputf("Authorized on account %s", b.bot.Self.UserName)
}

func main() {
	var testBot Bot
	testBot.Setup(os.Getenv("TeleBot"))
}