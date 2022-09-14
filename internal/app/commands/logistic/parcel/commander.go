package parcel

import (
	"encoding/json"
	"fmt"

	service "github.com/gempellm/bot/internal/service/logistic/parcel"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type ParcelCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message) // return error not implemented
	Edit(inputMsg *tgbotapi.Message)
}

type Commander struct {
	bot           *tgbotapi.BotAPI
	parcelService service.ParcelService
}

func NewParcelCommander(bot *tgbotapi.BotAPI, service service.ParcelService) *Commander {
	return &Commander{
		bot:           bot,
		parcelService: service,
	}
}

func (c *Commander) HandleUpdate(update tgbotapi.Update) {
	if update.CallbackQuery != nil {
		parsedData := CommandData{}
		json.Unmarshal([]byte(update.CallbackQuery.Data), &parsedData)

		parcels, _ := c.parcelService.List(parsedData.Offset, parsedData.Limit)

		msgText := fmt.Sprintf("Parcels from %d to %d:\n\n", parsedData.Offset, parsedData.Offset+parsedData.Limit)
		for _, parcel := range parcels {
			msgText += parcel.String()
			msgText += "\n"
		}

		msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, msgText)

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
		return
	}

	if update.Message != nil {
		switch update.Message.Command() {
		case "help", "start":
			c.Help(update.Message)
		case "get":
			c.Get(update.Message)
		case "list":
			c.List(update.Message)
		case "delete":
			c.Delete(update.Message)
		case "new":
			c.New(update.Message)
		}
	}
}
