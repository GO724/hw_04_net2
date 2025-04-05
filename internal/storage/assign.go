package storage

import "fmt"

// assign_$ID` - делаеть пользователя исполнителем задачи
func (s Stor) AssignTask() {
	fmt.Printf("Assign task %v", s)
}
