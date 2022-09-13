package commands

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Default(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text) // логируем сообщение

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text) // создаем ответ, указываем id чата и текст сообщения
	//msg.ReplyToMessageID = update.Message.MessageID                                       // указание того, что сообщение - это реплай на предыдущее сообщение

	c.bot.Send(msg)
}

func (c *Commander) HandleUpdate(update tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered from panic: %v", panicValue)
		}
	}()

	if update.Message != nil { // если получаем не пустое сообщение
		switch update.Message.Command() { // получаем команду из сообщения
		case "help", "start": // помощь, справочная информация
			c.Help(update.Message)
		case "list": // вывести список продуктов
			c.List(update.Message)
		case "get": // получить информацию о продукте
			c.Get(update.Message)
		default: // сообщение без команды
			c.Default(update.Message)
		}
	}
}
