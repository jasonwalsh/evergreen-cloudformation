package command

import (
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

type DeleteCommand struct {
	client *cloudformation.CloudFormation
	input  *cloudformation.DeleteStackInput
}

func DeleteStack(client *cloudformation.CloudFormation) Command {
	var input cloudformation.DeleteStackInput
	return &DeleteCommand{client, &input}
}

func (c *DeleteCommand) BuildRequest() {
	c.input.SetStackName("")
}

func (c *DeleteCommand) WaitUntilFinished() error {
	var input cloudformation.DescribeStacksInput
	input.SetStackName(*c.input.StackName)
	if err := c.client.WaitUntilStackDeleteComplete(&input); err != nil {
		return err
	}
	return nil
}

func (c *DeleteCommand) SendSync() error {
	if _, err := c.client.DeleteStack(c.input); err != nil {
		return err
	}
	if err := c.WaitUntilFinished(); err != nil {
		return err
	}
	return nil
}

func (c *DeleteCommand) Execute() error {
	c.BuildRequest()
	if err := c.SendSync(); err != nil {
		return err
	}
	return nil
}

func (c *DeleteCommand) SetParameters(v map[string]string) {
	return
}
