package storage

import "fmt"

// owner` - показывает задачи, которые были созданы мной
func (s Stor) OwnerTask() {
	fmt.Printf("Owner task %v", s)
}
