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

func (c *Commander) HandleUpdate(update tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recoverd from panic: %v", panicValue)
		}
	}()

	if update.CallbackQuery != nil {
		msg := tgbotapi.NewMessage(
			update.CallbackQuery.Message.Chat.ID,
			"Data:"+update.CallbackQuery.Data,
		)
		c.bot.Send(msg)
		return
	}

	if update.Message == nil {
		return
	}
	switch update.Message.Command() {
	case "help":
		c.Help(update.Message)
	case "get":
		c.Get(update.Message)
	case "list":
		c.List(update.Message)
	case "delete":
		c.Delete(update.Message)
	case "create":
		c.Create(update.Message)
	case "replace":
		c.Replace(update.Message)
	default:
		c.Default(update.Message)
	}
}
