package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

func (c *Commander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)

	if err != nil {
		log.Println("Wrong args ", args)
		return
	}

	product, err := c.productService.Get(idx)

	if err != nil {
		log.Printf("Fail To Product %v - %s", idx, err)
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		product.Title,
	)
	//inputMessage.ReplyToMessageID = inputMessage.MessageID

	c.bot.Send(msg)
}
