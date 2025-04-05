package controller

// Слой между внешним api (tgbot http handlers) и логикой (service)
// Функции структуры Controller вызываются как handler'ы http сервера tgbot api
// Доступ к структурам данных приложения через поля структуры Controller
// Внедрение зависимостей через конструктор New()

import (
	"hw_04_net2/internal/service"
)

type Controller struct {
	Service *service.Service
}

func New(s *service.Service) *Controller {
	return &Controller{
		Service: s,
	}
}
