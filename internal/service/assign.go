package service

import "fmt"

// assign_$ID` - делаеть пользователя исполнителем задачи
func (s Service) AssignTask() {
	fmt.Printf("Assign task %v", s)
}
