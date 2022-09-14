package parcel

import (
	"fmt"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMsg *tgbotapi.Message) {
	rawArgs := inputMsg.CommandArguments()
	args := strings.Split(rawArgs, " ")
	offset, err := strconv.ParseUint(args[0], 10, 64)
	if err != nil {
		msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "Wrong input format.")
		c.bot.Send(msg)
		return
	}

	limit, err := strconv.ParseUint(args[1], 10, 64)
	if err != nil {
		msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "Wrong input format.")
		c.bot.Send(msg)
		return
	}

	parcels, _ := c.parcelService.List(offset, limit)

	msgText := fmt.Sprintf("Parcels from %d to %d:\n\n", offset, offset+limit)
	for _, parcel := range parcels {
		msgText += parcel.String()
		msgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, msgText)
	c.bot.Send(msg)
}
