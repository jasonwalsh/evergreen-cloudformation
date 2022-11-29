package main

import (
	"flag"
	"log"

	"github.com/jasonwalsh/cloudformation/command"
)

var action command.Action
var project string

func init() {
	flag.Var(&action, "action", "")
	flag.StringVar(&project, "project", "", "")
}

func main() {
	flag.Parse()
	command := command.New(action)
	command.SetParameters(map[string]string{
		"BucketName": project,
	})
	if err := command.Execute(); err != nil {
		log.Fatal(err)
	}
}
