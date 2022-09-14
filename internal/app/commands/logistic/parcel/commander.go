package parcel

import (
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
