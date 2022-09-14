package parcel

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Help(inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprint(
		"/help - List all available commands\n",
		"/get parcel_ID - Get info about parcel\n",
		"/list index n - Get n parcels from specified index\n",
		"/delete parcel_ID - Delete parcel with specified parcelID\n",
		"/new title - Create parcel with specified title\n",
		"/edit parcel_ID title - Edit parcel's title\n",
	))

	c.bot.Send(msg)
}
