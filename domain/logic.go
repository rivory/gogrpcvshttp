package domain

import "strings"

// LogicInterface implements the logic service
type LogicInterface interface {
	Uppercase(message string) string
}

// LogicService implements the logic layer
type LogicService struct{}

// Uppercase does take a string as parameters, and return it in uppercase
func (ls LogicService) Uppercase(message string) string {
	return strings.ToUpper(message)
}
