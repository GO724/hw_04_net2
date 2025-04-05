package service

import "fmt"

func (s Service) Start() {
	fmt.Printf("Unassign task %v", s)
}
