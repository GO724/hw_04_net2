package controller

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (c *Controller) UnassignTaskHandler(ctx context.Context, b *bot.Bot, update *models.Update) {

	// если задача назначена и не на мне - показывается логин исполнителя :331
	// /unassign_ - снимает задачу с себя, нельзя снять задачу которая не на вас :343 (output `Задача не на вас`)
	// /unassign_ - снимает задачу с себя, автору отправляется уведомление о том, что задача осталась без исполнителя :352
	// Petrov,
	// "/unassign_1",
	// map[int64]string{
	// 	Petrov: `Принято`,
	// 	Ivanov: `Задача "написать бота" осталась без исполнителя`

}
