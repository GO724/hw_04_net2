package service

import (
	"context"
	"fmt"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// my` - показывает задачи, которые назначены на меня
func (s Service) MyTask(ctx context.Context, b *bot.Bot, update *models.Update) {
	fmt.Printf("My task %v", s)
}
