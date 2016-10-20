package main

import (
	"fmt"

	"github.com/opencontainers/runc/libcontainer"
	"github.com/urfave/cli"
)

var attachCommand = cli.Command{
	Name:  "attach",
	Usage: "attach I/O stream to specified container",
	ArgsUsage: `<container-id>

Where "<container-id>" is your name for the instance of the container that you
started before. The name you provide for the container instance must be unique 
on your host, and the container should be in running state.`,
	Description: `The attach command allow you to attach I/O stream to one running container.`,
	Action: func(context *cli.Context) error {
		container, err := getContainer(context)
		if err != nil {
			return err
		}
		status, err := container.Status()
		if err != nil {
			return err
		}
		switch status {
		case libcontainer.Running:
			return container.Exec()
		default:
			return fmt.Errorf("cannot attach to a container in the %s state", status)
		}
	},
}
