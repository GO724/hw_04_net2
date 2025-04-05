package storage

import "fmt"

// unassign_$ID` - снимает задачу с текущего исполнителя
func (s Stor) UnassignTask() {
	fmt.Printf("Unassign task %v", s)
}
