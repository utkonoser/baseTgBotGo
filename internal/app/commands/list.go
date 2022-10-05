package commands

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	outputMsg := "Here all the products:\n\n"
	products := c.productService.List()
	for _, item := range products {
		outputMsg += item.Title +"\n"
	}

	outputIndexes := "\n"
	for idx, item := range products {
		outputIndexes += fmt.Sprintf("%s - %v\n",item.Title,idx)
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsg)

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Show indexes", outputIndexes),
		),
	)

	c.bot.Send(msg)
}
