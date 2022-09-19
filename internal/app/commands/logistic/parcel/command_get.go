package parcel

import (
	"strconv"

	"github.com/gempellm/bot/internal/model/logistic"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Get(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()
	parcelID, err := strconv.ParseUint(args, 10, 64)

	if err != nil {
		msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "Wrong input format.")
		c.bot.Send(msg)
		return
	}

	parcel, err := c.parcelService.Describe(parcelID)
	if err != nil {
		msg := tgbotapi.NewMessage(inputMsg.Chat.ID, logistic.StringParcelNotFound(parcelID))
		c.bot.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, parcel.String())
	c.bot.Send(msg)
}
