package actions

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func Init(app *cli.App) {
	action := cli.Command{
		Name:  "init",
		Usage: "create a niflheim environment file",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "force",
				Aliases: []string{"f"},
				Usage:   "Force overwrite of existing niflheim.env",
			},
		},
		Action: func(ctx *cli.Context) error {
			envfile, err := envFilename()
			if err != nil {
				return fmt.Errorf("getting env filename: %w", err)
			}
			if !ctx.Bool("force") {
				if fileExists(envfile) {
					return fmt.Errorf("file %s already exists, use -f to overwrite", envfile)
				}
			}
			data, err := stdData()
			if err != nil {
				return fmt.Errorf("getting std data: %w", err)
			}
			out, err := render(initTemplate, data)
			if err != nil {
				return fmt.Errorf("rendering initTemplate: %w", err)
			}

			err = os.WriteFile(envfile, []byte(out), 0644)
			if err != nil {
				return fmt.Errorf("writing %s: %w", envfile, err)
			}
			fmt.Println("\nCreated config file:", envfile)
			fmt.Printf("\nPlease check the contents of %s and edit it accordingly\n", envfile)
			return nil
		},
	}
	app.Commands = append(app.Commands, &action)
}

var initTemplate = `VALHEIM_SERVER_FOLDER={{.HOME}}/valheim-server
VALHEIM_DATA_FOLDER={{.HOME}}/valheim-data
VALHEIM_LOGS_FOLDER={{.HOME}}/valheim-logs
VALHEIM_START_SCRIPT_NAME=niflheim-start.sh

VALHEIM_SERVER_NAME="Niflheim"
VALHEIM_SERVER_PASSWORD="h31nl31n"
VALHEIM_SERVER_PORT=2456
VALHEIM_SERVER_WORLD="niflheim"
VALHEIM_SERVER_PUBLIC=1

VALHEIM_SERVICE_NAME=valheim-{{.USER}}
`
