package commands

import (
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/jasonwalsh/cloudformation/commands/create"
	"github.com/jasonwalsh/cloudformation/commands/delete"
)

type Command interface {
	BuildRequest()
	Execute() error
	SendSync() error
	SetParameters(map[string]string)
	WaitUntilFinished() error
}

func CreateStack(cloudformation *cloudformation.CloudFormation) Command {
	return create.New(cloudformation)
}

func DeleteStack(cloudformation *cloudformation.CloudFormation) Command {
	return delete.New(cloudformation)
}
