package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	//"z1/internal/service/product"
)

func (c *Commander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/HELP - HELP\n"+
			"/LIST - LIST",
	)
	c.bot.Send(msg)
}
