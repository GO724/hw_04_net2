package service

import "fmt"

// unassign_$ID` - снимает задачу с текущего исполнителя
func (s Service) UnassignTask() {
	fmt.Printf("Unassign task %v", s)
}
