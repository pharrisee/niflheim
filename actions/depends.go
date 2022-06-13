package actions

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func Depends(app *cli.App) {
	action := cli.Command{
		Name:  "depends",
		Usage: "install required dependencies",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "update",
				Usage: "fully update OS packages",
				Aliases: []string{
					`u`,
				},
			},
		},
		Action: func(_ *cli.Context) error {
			if err := LoadEnv(); err != nil {
				return fmt.Errorf("failed to load environment, use 'niflheim init' to create a skeleton environment file: %s", err)
			}

			for _, command := range dependsCommands {
				if err := runCmd(command); err != nil {
					return fmt.Errorf("%s : %v", command, err)
				}
			}

			fmt.Println("\nDependency installation complete.")

			return nil
		},
	}

	app.Commands = append(app.Commands, &action)
}

var dependsCommands = []string{
	`sudo apt-get update`,
	`sudo apt-get install --no-install-recommends --no-install-suggests -y software-properties-common`,
	`sudo dpkg --add-architecture i386`,
	`sudo add-apt-repository multiverse`,
	`echo PURGE | sudo debconf-communicate steam`,
	`echo PURGE | sudo debconf-communicate steamcmd`,
	`echo steam steam/question select "I AGREE" | sudo debconf-set-selections`,
	`echo steam steam/license note '' | sudo debconf-set-selections`,
	`echo steam steam/purge note '' | sudo debconf-set-selections`,
	`sudo apt-get install --no-install-recommends --no-install-suggests -y steamcmd lib32gcc-s1 lib32stdc++6 libsdl2-2.0-0:i386 libsdl2-2.0-0`,
	`/usr/games/steamcmd +quit &> /dev/null`,
	`sudo apt-get full-upgrade -y --allow-downgrades`,
	`sudo apt-get autoremove -y`,
	`sudo apt-get autoclean -y`,
}
