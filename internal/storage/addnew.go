package storage

import "fmt"

// /new XXX YYY ZZZ` - создаёт новую задачу
func (s Stor) NewTask() {
	fmt.Printf("New task %v", s)
}
