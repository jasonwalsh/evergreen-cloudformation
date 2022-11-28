package commands

import (
	"fmt"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

var templateBody = `
{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Parameters": {
    "BucketName": {
      "Type": "String"
    }
  },
  "Resources": {
    "AccessKey": {
      "Properties": {
        "UserName": {
          "Ref": "User"
        }
      },
      "Type": "AWS::IAM::AccessKey"
    },
    "Bucket": {
      "Properties": {
        "BucketName": {
          "Ref": "BucketName"
        }
      },
      "Type": "AWS::S3::Bucket"
    },
    "BucketPolicy": {
      "Properties": {
        "Bucket": {
          "Ref": "Bucket"
        },
        "PolicyDocument": {
          "Statement": [
            {
              "Action": ["s3:*"],
              "Effect": "Allow",
              "Principal": {
                "AWS": {
                  "Fn::GetAtt": ["User", "Arn"]
                }
              },
              "Resource": {
                "Fn::GetAtt": ["Bucket", "Arn"]
              }
            }
          ],
          "Version": "2012-10-17"
        }
      },
      "Type": "AWS::S3::BucketPolicy"
    },
    "User": {
      "Properties": {
        "UserName": {
          "Ref": "Bucket"
        }
      },
      "Type": "AWS::IAM::User"
    }
  }
}
`

type CreateCommand struct {
	client *cloudformation.CloudFormation
	input  *cloudformation.CreateStackInput
}

func CreateStack(client *cloudformation.CloudFormation) Command {
	var input cloudformation.CreateStackInput
	return &CreateCommand{client, &input}
}

func UniqueName() string {
	return fmt.Sprintf("id-%s", strconv.FormatInt(time.Now().UnixNano(), 10))
}

func (c *CreateCommand) BuildRequest() {
	c.input.SetCapabilities([]*string{aws.String(cloudformation.CapabilityCapabilityNamedIam)})
	c.input.SetEnableTerminationProtection(false)
	c.input.SetStackName(UniqueName())
	c.input.SetTemplateBody(templateBody)
}

func (c *CreateCommand) WaitUntilFinished() error {
	var input cloudformation.DescribeStacksInput
	input.SetStackName(*c.input.StackName)
	if err := c.client.WaitUntilStackCreateComplete(&input); err != nil {
		return err
	}
	return nil
}

func (c *CreateCommand) SendSync() error {
	if _, err := c.client.CreateStack(c.input); err != nil {
		return err
	}
	if err := c.WaitUntilFinished(); err != nil {
		return err
	}
	return nil
}

func (c *CreateCommand) Execute() error {
	c.BuildRequest()
	if err := c.SendSync(); err != nil {
		return err
	}
	return nil
}

func (c *CreateCommand) SetParameters(v map[string]string) {
	parameters := make([]*cloudformation.Parameter, len(v))
	for key, value := range v {
		parameter := cloudformation.Parameter{
			ParameterKey:   aws.String(key),
			ParameterValue: aws.String(value),
		}
		parameters = append(parameters, &parameter)
	}
	c.input.SetParameters(parameters)
}
