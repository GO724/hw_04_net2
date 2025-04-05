package service

import "fmt"

// owner` - показывает задачи, которые были созданы мной
func (s Service) OwnerTask() {
	fmt.Printf("Owner task %v", s)
}
