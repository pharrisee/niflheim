package actions

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func ServiceRestart(app *cli.App) {
	init := cli.Command{
		Name:  "service-restart",
		Usage: "start the service for the instance",
		Action: func(*cli.Context) error {
			if err := LoadEnv(); err != nil {
				return fmt.Errorf("failed to load environment, use 'niflheim init' to create a skeleton environment file: %s", err)
			}
			if err := runCmd("sudo systemctl restart " + os.Getenv("VALHEIM_SERVICE_NAME")); err != nil {
				return fmt.Errorf("failed to restart %s: %s", os.Getenv("VALHEIM_SERVICE_NAME"), err)
			}
			if err := runCmd("sudo systemctl status --no-pager " + os.Getenv("VALHEIM_SERVICE_NAME")); err != nil {
				return fmt.Errorf("failed to get status of %s: %s", os.Getenv("VALHEIM_SERVICE_NAME"), err)
			}
			return nil
		},
	}
	app.Commands = append(app.Commands, &init)
}
