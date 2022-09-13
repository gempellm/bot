package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const token = "TOKEN"

func main() {
	bot, err := tgbotapi.NewBotAPI(token) // создаем объект бота
	if err != nil {
		log.Panic(err) // если не вышло - падаем
	}

	bot.Debug = true // включаем режим дебага, который выводит все сообщения, которые пришли к боту

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	} // инициализируются настройки подписки на сообщения

	updates := bot.GetUpdatesChan(u) // получаем канал (очередь) сообщений

	for update := range updates {
		if update.Message != nil { // если получаем не пустое сообщение
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text) // логируем сообщение

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You wrote: "+update.Message.Text) // создаем ответ, указываем id чата и текст сообщения
			msg.ReplyToMessageID = update.Message.MessageID                                       // указание того, что сообщение - это реплай на предыдущее сообщение

			bot.Send(msg)
		}
	}
}
