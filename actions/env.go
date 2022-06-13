package actions

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func Env(app *cli.App) {
	action := cli.Command{
		Name:  "env",
		Usage: "print the current environment",
		Action: func(*cli.Context) error {
			if err := LoadEnv(); err != nil {
				return fmt.Errorf("failed to load environment, use 'niflheim init' to create a skeleton environment file: %s", err)
			}
			env := os.Environ()
			for _, v := range env {
				fmt.Println(v)
			}
			return nil
		},
	}
	app.Commands = append(app.Commands, &action)
}
