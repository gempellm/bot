package commands

import (
	"github.com/gempellm/bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(
	bot *tgbotapi.BotAPI,
	productService *product.Service,
) *Commander {
	return &Commander{
		bot:            bot,
		productService: productService,
	}
}

func (c *Commander) HandleUpdate(update tgbotapi.Update) {
	if update.Message != nil { // если получаем не пустое сообщение
		switch update.Message.Command() { // получаем команду из сообщения
		case "help", "start": // помощь, справочная информация
			c.Help(update.Message)
		case "list":
			c.List(update.Message)
		case "get":
			c.Get(update.Message)
		default: // сообщение без команды
			c.Default(update.Message)
		}
	}
}
