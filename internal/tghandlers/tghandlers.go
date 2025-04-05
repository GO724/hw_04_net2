package tghandlers

// Внешнее api

import (
	"hw_04_net2/internal/controller"

	"github.com/go-telegram/bot"
)

type TgHandlers struct {
	Bot        *bot.Bot
	Controller *controller.Controller
}

func New(bot *bot.Bot, controller *controller.Controller) *TgHandlers {
	return &TgHandlers{
		Bot:        bot,
		Controller: controller,
	}
}

func (th *TgHandlers) Bindle() map[string]string {
	// Управление происходит через текстовый интерфейс:
	// /tasks
	// /new XXX YYY ZZZ - создаёт новую задачу
	// /assign_$ID - делаеть пользователя исполнителем задачи
	// /unassign_$ID - снимает задачу с текущего исполнителя
	// /resolve_$ID - выполняет задачу, удаляет её из списка
	// /my - показывает задачи, которые назначены на меня
	// /owner - показывает задачи, которые были созданы мной Подробности форматирования смотрите в тестах.

	// Регистрация обработчиков контроллера, через который внедряем зависимости
	handlersID := make(map[string]string, 5)

	b := th.Bot
	c := th.Controller
	handlersID["/"] = b.RegisterHandler(bot.HandlerTypeMessageText, "/", bot.MatchTypePrefix, c.StartHandler)
	handlersID["/new"] = b.RegisterHandler(bot.HandlerTypeMessageText, "/new", bot.MatchTypePrefix, c.NewTaskHandler)
	handlersID["/assign_"] = b.RegisterHandler(bot.HandlerTypeMessageText, "/assign_", bot.MatchTypePrefix, c.AssignTaskHandler)
	handlersID["/unassign_"] = b.RegisterHandler(bot.HandlerTypeMessageText, "/unassign_", bot.MatchTypePrefix, c.UnassignTaskHandler)
	handlersID["/resolve_"] = b.RegisterHandler(bot.HandlerTypeMessageText, "/resolve_", bot.MatchTypePrefix, c.ResolveTaskHandler)

	return handlersID
}
