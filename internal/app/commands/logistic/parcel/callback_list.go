package parcel

import (
	"encoding/json"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CommandData struct {
	Command string
	Offset  uint64
	Limit   uint64
}

func (c *Commander) CallbackList(callback *tgbotapi.CallbackQuery) {
	parsedData := CommandData{}
	json.Unmarshal([]byte(callback.Data), &parsedData)

	parcels, _ := c.parcelService.List(parsedData.Offset, parsedData.Limit)

	msgText := fmt.Sprintf("Parcels from %d to %d:\n\n", parsedData.Offset, parsedData.Offset+parsedData.Limit)
	for _, parcel := range parcels {
		msgText += parcel.String()
		msgText += "\n"
	}

	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, msgText)

	if uint64(len(parcels)) == parsedData.Limit {
		serializedData, _ := json.Marshal(CommandData{
			Command: "list",
			Offset:  parsedData.Offset + parsedData.Limit,
			Limit:   parsedData.Limit,
		})

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Next page", string(serializedData)),
			),
		)
	}

	c.bot.Send(msg)
}
