package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "/help - help\n"+
		"/list - list products\n"+
		"/get [index] - get product for index\n"+
		"/delete [index] - delete product for index\n"+
		"/create [product] - for create new product\n"+
		"/replace [index][space][args] - replace product for index",
	)

	c.bot.Send(msg)
}
