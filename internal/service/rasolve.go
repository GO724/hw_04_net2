package service

import "fmt"

// resolve_$ID` - выполняет задачу, удаляет её из списка
func (s Service) ResolveTask() {
	fmt.Printf("Resolve task %v", s)
}
