package actions

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func ServiceInstall(app *cli.App) {
	action := cli.Command{
		Name:  "service-install",
		Usage: "install the service for the instance",
		Action: func(*cli.Context) error {
			if err := LoadEnv(); err != nil {
				return fmt.Errorf("failed to load environment, use 'niflheim init' to create a skeleton environment file: %s", err)
			}
			data, err := stdData()
			if err != nil {
				return fmt.Errorf("getting standard data: %w", err)
			}

			serviceFileTmp := fmt.Sprintf("/tmp/%s.service", os.Getenv("VALHEIM_SERVICE_NAME"))
			serviceFile := fmt.Sprintf("/etc/systemd/system/%s.service", os.Getenv("VALHEIM_SERVICE_NAME"))

			contents, err := render(unitFileTemplate, data)
			if err != nil {
				return fmt.Errorf("failed to fill unitFileTemplate: %s", err)
			}
			if err := os.WriteFile(serviceFileTmp, []byte(contents), 0644); err != nil {
				return fmt.Errorf("failed to write unit file (%s): %s", serviceFileTmp, err)
			}

			if err := runCmd(fmt.Sprintf("/bin/sh -c 'sudo cp %s %s'", serviceFileTmp, serviceFile)); err != nil {
				return fmt.Errorf("failed to copy unit file (%s -> %s): %s", serviceFileTmp, serviceFile, err)
			}

			if err := runCmd(fmt.Sprintf("sudo systemctl enable %s", os.Getenv("VALHEIM_SERVICE_NAME"))); err != nil {
				return fmt.Errorf("failed to enable seven: %s", err)
			}

			if err := runCmd("sudo systemctl daemon-reload"); err != nil {
				return fmt.Errorf("failed to reload daemon: %s", err)
			}
			return nil

		},
	}
	app.Commands = append(app.Commands, &action)
}

var unitFileTemplate = `# unit file for Valheim installed at {{.VALHEIM_SERVER_FOLDER}}
[Unit]
Description={{.VALHEIM_SERVICE_NAME}}
After=network.target

[Service]
Type=simple
User={{.USER}}
WorkingDirectory={{.VALHEIM_SERVER_FOLDER}}
ExecStart={{.VALHEIM_SERVER_FOLDER}}/niflheim-start.sh
KillSignal=SIGINT
Restart=on-failure
TimeoutStopSec=10

[Install]
WantedBy=multi-user.target
`
