package entity

// Описание сущностей приложения

// tasks_test.go:89
// Ivanov: &tgbotapi.User{
// 	ID:           Ivanov,
// 	FirstName:    "Ivan",
// 	LastName:     "Ivanov",
// 	UserName:     "ivanov",
// 	LanguageCode: "ru",
// 	IsBot:        false,
// },

type User struct {
	UUID         int
	ID           string
	FirstName    string
	LastName     string
	UserName     string
	LanguageCode string
	IsBot        bool
}

type Task struct {
	ID       int
	OwnerID  int
	UserID   int
	Task     string
	Complete bool
}
