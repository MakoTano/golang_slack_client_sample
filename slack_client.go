package main

import (
	"fmt"
	"os"

	"github.com/MakoTano/golang_slack_client_sample/model"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "slack_client"
	app.Usage = "this posts a parameter string to slack using Webhook"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name:  "post",
			Usage: "post [text]",
			Action: func(c *cli.Context) {
				var s, err = model.NewSlack()
				if err != nil {
					fmt.Println(err.Error())
					return
				}

				// you can set any parameters here.
				if text := c.Args().First(); text != "" {
					s.Text = text
				}

				if err = s.Post(); err != nil {
					fmt.Println(err.Error())
					return
				}

				fmt.Println("Posted!")
				return
			},
		},
	}

	app.Run(os.Args)
}
