package commands

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Default(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)

	c.bot.Send(msg)
}

func (c *Commander) HandleUpdate(inputMessage *tgbotapi.Message) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recoverd from panic: %v", panicValue)
		}
	}()

	if inputMessage == nil {
		return
	}
	switch inputMessage.Command() {
	case "help":
		c.Help(inputMessage)
	case "get":
		c.Get(inputMessage)
	case "list":
		c.List(inputMessage)
	default:
		c.Default(inputMessage)
	}
}
