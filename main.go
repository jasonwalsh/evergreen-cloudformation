package main

import (
	"flag"
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/jasonwalsh/cloudformation/commands"
)

var delete bool
var project string

func init() {
	flag.BoolVar(&delete, "delete", false, "")
	flag.StringVar(&project, "project", "", "")
}

func NewSession() *session.Session {
	return session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
}

func main() {
	flag.Parse()
	cloudformation := cloudformation.New(NewSession())
	var command commands.Command
	if delete {
		command = commands.DeleteStack(cloudformation)
	} else {
		command = commands.CreateStack(cloudformation)
	}
	command.SetParameters(map[string]string{
		"BucketName": project,
	})
	if err := command.Execute(); err != nil {
		log.Fatal(err)
	}
}
