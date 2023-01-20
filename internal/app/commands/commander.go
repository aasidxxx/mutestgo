package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"z1/internal/service/product"
)

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI, productService *product.Service) *Commander {
	return &Commander{
		bot:            bot,
		productService: productService,
	}
}

//func (c *Commander) Help(inputMessage *tgbotapi.Message) {
//	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
//		"/HELP - HELP\n"+
//			"/LIST - LIST",
//	)
//	c.bot.Send(msg)
//}

//func (c *Commander) List(inputMessage *tgbotapi.Message) {
//
//	outputMsg := "Herer all product: \n\n"
//
//	product := c.productService.List()
//
//	for _, p := range product {
//		outputMsg += p.Title
//		outputMsg += "\n"
//	}
//
//	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsg)
//	c.bot.Send(msg)
//}

//func (c *Commander) Default(inputMessage *tgbotapi.Message) {
//	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, inputMessage.Text)
//	//inputMessage.ReplyToMessageID = inputMessage.MessageID
//
//	c.bot.Send(msg)
//}
