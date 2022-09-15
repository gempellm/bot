package parcel

import (
	"fmt"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Delete(inputMsg *tgbotapi.Message) {
	rawData := inputMsg.CommandArguments()

	args := strings.Split(rawData, " ")

	for _, arg := range args {
		c.removeArg(arg, inputMsg)
	}
}

func (c *Commander) removeArg(arg string, inputMsg *tgbotapi.Message) {
	id, err := strconv.ParseUint(arg, 10, 64)

	if err != nil {
		msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "Wrong input format.")
		c.bot.Send(msg)
		return
	}

	ok, _ := c.parcelService.Remove(id)
	if !ok {
		msg := tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprintf("Parcel with ID %d not found.", id))
		c.bot.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprintf("Parcel with ID %d was successfully removed.", id))
	c.bot.Send(msg)
}
