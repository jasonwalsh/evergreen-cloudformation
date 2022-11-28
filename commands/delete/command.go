package delete

import (
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

type command struct {
	client *cloudformation.CloudFormation
	input  *cloudformation.DeleteStackInput
}

func New(client *cloudformation.CloudFormation) *command {
	var input cloudformation.DeleteStackInput
	return &command{client, &input}
}

func (c *command) BuildRequest() {
	c.input.SetStackName("")
}

func (c *command) WaitUntilFinished() error {
	var input cloudformation.DescribeStacksInput
	input.SetStackName(*c.input.StackName)
	if err := c.client.WaitUntilStackDeleteComplete(&input); err != nil {
		return err
	}
	return nil
}

func (c *command) SendSync() error {
	if _, err := c.client.DeleteStack(c.input); err != nil {
		return err
	}
	if err := c.WaitUntilFinished(); err != nil {
		return err
	}
	return nil
}

func (c *command) Execute() error {
	c.BuildRequest()
	if err := c.SendSync(); err != nil {
		return err
	}
	return nil
}

func (c *command) SetParameters(v map[string]string) {
	return
}
