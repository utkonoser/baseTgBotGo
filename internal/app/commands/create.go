package commands

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Create(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	if args == "" {
		log.Printf("fail, args is empty")
		return
	}

	product, err := c.productService.Create(args)
	if err != nil {
		log.Printf("fail to create product %s:%v", product, err)
		return
	}

	outputMsg := "Successfully add a new product!\nThis is new list:\n"
	products := c.productService.List()
	for _, item := range products {
		outputMsg += item.Title + "\n"
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		outputMsg)

	c.bot.Send(msg)
}
