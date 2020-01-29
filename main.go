package main

import (
	"fmt"
	"os"

	"github.com/lunarxlark/cfn-snippet/cmd"
	"github.com/urfave/cli"
)

func main() {
	err := newApp().Run(os.Args)
	if err != nil {
		fmt.Println("fatal")
	}
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "test"
	app.Usage = "test usage"
	app.Version = "0.0.1"
	app.Commands = cmd.Commands

	return app
}
