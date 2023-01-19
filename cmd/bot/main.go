package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
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

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			switch update.Message.Command() {
			case "help":
				helpCommand(bot, update.Message)
			case "list":
				listCommand(bot, update.Message, productService)
			default:
				defaultBehavior(bot, update.Message)
			}

		}
	}
}

func helpCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/HELP - HELP\n"+
			"/LIST - LIST",
	)
	bot.Send(msg)
}

func listCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message, productService product.Service) {

	outputMsg := "Herer all product: \n\n"

	product := productService.List()

	for _, p := range product {
		outputMsg += p.Title
		outputMsg += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsg)
	bot.Send(msg)
}

func defaultBehavior(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, inputMessage.Text)
	//inputMessage.ReplyToMessageID = inputMessage.MessageID

	bot.Send(msg)
}
