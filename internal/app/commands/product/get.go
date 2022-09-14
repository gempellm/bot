package commands

import (
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	idx, err := strconv.Atoi(args)
	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Wrong argument.")
		c.bot.Send(msg)
		return
	}

	product, err := c.productService.Get(idx)
	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Product not found.")
		c.bot.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		product.Title,
	)
	c.bot.Send(msg)
}
