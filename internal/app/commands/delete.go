package commands

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args: ", args)
		return
	}

	_, err = c.productService.Delete(idx)
	if err != nil {
		log.Printf("fail to get profuct with idx %d:%v", idx, err)
		return
	}

	outputMsg := "This is new list:\n"
	products := c.productService.List()
	for _, item := range products {
		outputMsg += item.Title + "\n"
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		outputMsg)

	c.bot.Send(msg)
}
