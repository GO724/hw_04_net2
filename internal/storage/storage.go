package storage

// Хранилище, к которому обращается service

import "hw_04_net2/internal/entity"

type Stor struct {
	Users map[int]entity.User
	Task  map[int]entity.Task
}

func New() (*Stor, error) {
	return &Stor{
		Users: make(map[int]entity.User, 100),
		Task:  make(map[int]entity.Task, 100),
	}, nil
}
