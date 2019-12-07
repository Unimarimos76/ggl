package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"os/exec"
)

func main() {
	app := cli.NewApp()
	app.Name = "sampleApp"
	app.Usage = "This app echo input arguments"
	app.Version = "0.0.1"

	app.Action = func(context *cli.Context) error {
		err := exec.Command("open", "https://www.google.com/search?q="+context.Args().Get(0)).Start()
		if err != nil {
			panic(err)
		}
		return nil
	}

	app.Run(os.Args)
}
