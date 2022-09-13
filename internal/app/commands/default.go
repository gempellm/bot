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
