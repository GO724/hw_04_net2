package service

import "hw_04_net2/internal/storage"

// Логика. Вызывается в controller зависит от storage

type Service struct {
	DB *storage.Stor
}

func New(db *storage.Stor) *Service {
	return &Service{
		DB: db,
	}
}
