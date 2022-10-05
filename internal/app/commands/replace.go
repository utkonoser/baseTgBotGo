package commands

import (
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Replace(inputMessage *tgbotapi.Message) {
	args := strings.Split(inputMessage.CommandArguments(), "")
	index := ""
	for _, arg := range args {
		if arg == " " {
			break
		}
		index += arg
	}

	idx, err := strconv.Atoi(index)
	if err != nil {
		log.Println("wrong args: ", args)
		return
	}

	newArg := strings.TrimSpace(strings.Join(args[1:], ""))

	_, err = c.productService.Replace(idx, newArg)
	if err != nil {
		log.Printf("fail to replace profuct with idx %d:%v", idx, err)
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
