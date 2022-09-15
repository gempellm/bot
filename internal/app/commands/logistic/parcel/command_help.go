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
		"/list_ids - Get list of all parcel ids\n",
		"/delete parcel_ID1 parcel_ID2... - Delete parcel(s) with specified parcel_ID\n",
		"/new title1 title2... - Create parcel(s) with specified title\n",
		"/edit parcel_ID title - Edit parcel's title\n",
	))

	c.bot.Send(msg)
}
