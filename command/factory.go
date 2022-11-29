package command

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

type Command interface {
	BuildRequest()
	Execute() error
	SendSync() error
	SetParameters(map[string]string)
	WaitUntilFinished() error
}

var commands = map[Action]func(client *cloudformation.CloudFormation) Command{
	CreateAction: CreateStack,
	DeleteAction: DeleteStack,
}

func New(action Action) Command {
	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	if command, exists := commands[action]; exists {
		return command(cloudformation.New(session))
	}
	return nil
}
