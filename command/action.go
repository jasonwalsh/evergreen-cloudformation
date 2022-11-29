package command

import (
	"fmt"
)

type Action string

const (
	CreateAction Action = "create"
	DeleteAction Action = "delete"
)

func (a *Action) Set(value string) error {
	switch value {
	case "create":
		*a = CreateAction
	case "delete":
		*a = DeleteAction
	default:
		*a = CreateAction
	}
	return nil
}

func (a *Action) String() string {
	return fmt.Sprint(*a)
}
