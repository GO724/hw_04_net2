package storage

import "fmt"

// my` - показывает задачи, которые назначены на меня
func (s Stor) MyTask() {
	fmt.Printf("My task %v", s)
}
