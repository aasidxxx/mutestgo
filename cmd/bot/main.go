package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"z1/internal/app/commands"
	"z1/internal/service/product"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var token = ""

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("TOKEN")
	//-----------------

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	productService := product.NewService()

	commander := commands.NewCommander(bot, productService)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			switch update.Message.Command() {
			case "help":
				commander.Help(update.Message)
			case "list":
				commander.List(update.Message)
			default:
				commander.Default(update.Message)
			}

		}
	}
}

//type Commander struct {
//	bot *tgbotapi.BotAPI
//	productService *product.Service
//}
//
//func NewCommander(bot *tgbotapi.BotAPI, productService *product.Service) *Commander {
//	return &Commander{
//		bot: bot,
//		productService: productService,
//	}
//}
//
//func (c *Commander) Help(inputMessage *tgbotapi.Message) {
//	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
//		"/HELP - HELP\n"+
//			"/LIST - LIST",
//	)
//	c.bot.Send(msg)
//}
//
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
//
//func (c *Commander)  Default(inputMessage *tgbotapi.Message) {
//	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, inputMessage.Text)
//	//inputMessage.ReplyToMessageID = inputMessage.MessageID
//
//	c.bot.Send(msg)
//}
