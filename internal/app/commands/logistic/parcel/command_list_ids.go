package parcel

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) ListIds(inputMsg *tgbotapi.Message) {
	ids, _ := c.parcelService.Ids()
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprint(ids))
	c.bot.Send(msg)
}
