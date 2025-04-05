package controller

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (c *Controller) MyTaskHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	// если задача назначена и на мне - показывается "на меня" tasks_test.go:321
	c.Service.MyTask(ctx, b, update)
}
