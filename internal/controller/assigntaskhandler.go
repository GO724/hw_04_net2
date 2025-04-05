package controller

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (c *Controller) AssignTaskHandler(ctx context.Context, b *bot.Bot, update *models.Update) {

	// /assign_* - назначает задачу на себя tasks_test.go:302
	// если задача назначена и на мне - показывается "на меня"

	// в случае если задача была назначена на кого-то - он получает уведомление об этом tasks_test.go:311
	// в данном случае она была назначена на Alexandrov, поэтому ему отправляется уведомление tasks_test.go:312

}
