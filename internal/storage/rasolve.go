package storage

import "fmt"

// resolve_$ID` - выполняет задачу, удаляет её из списка
func (s Stor) ResolveTask() {
	fmt.Printf("Resolve task %v", s)
}
