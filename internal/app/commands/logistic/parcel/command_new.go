package parcel

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gempellm/bot/internal/model/logistic"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) New(inputMsg *tgbotapi.Message) {
	var lastID uint64
	data, _ := os.ReadFile("lastID.txt")
	lastID, _ = strconv.ParseUint(string(data), 10, 64)
	lastID++
	os.WriteFile("lastID.txt", []byte(fmt.Sprint(lastID)), 0666)

	title := inputMsg.CommandArguments()

	parcel := logistic.Parcel{Title: title, ParcelID: lastID}

	parcelID, err := c.parcelService.Create(parcel)
	if err != nil {
		msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "Error occured during parcel creation.")
		c.bot.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprintf("Successfully created parcel with parcelID = %d", parcelID))
	c.bot.Send(msg)
}
