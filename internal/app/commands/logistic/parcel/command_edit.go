package parcel

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gempellm/bot/internal/model/logistic"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Edit(inputMsg *tgbotapi.Message) {
	rawArgs := inputMsg.CommandArguments()
	args := strings.Split(rawArgs, " ")
	if len(args) != 2 {
		msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "Wrong input format.")
		c.bot.Send(msg)
		return
	}

	parcelID, err := strconv.ParseUint(args[0], 10, 64)
	if err != nil {
		msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "Wrong input format.")
		c.bot.Send(msg)
		return
	}

	title := args[1]

	parcel, err := c.parcelService.Describe(parcelID)
	if err != nil {
		msg := tgbotapi.NewMessage(inputMsg.Chat.ID, logistic.StringParcelNotFound(pracelID))
		c.bot.Send(msg)
		return
	}

	parcel.Title = title

	c.parcelService.Update(parcelID, *parcel)
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprintf("Parcel with ID %d was successfully updated.", parcelID))
	c.bot.Send(msg)
}
