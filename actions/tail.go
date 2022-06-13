package actions

import (
	"fmt"
	"os"
	"sort"

	"github.com/hpcloud/tail"
	"github.com/urfave/cli/v2"
)

func Tail(app *cli.App) {
	init := cli.Command{
		Name:  "tail",
		Usage: "tail the log",
		Action: func(*cli.Context) error {
			if err := LoadEnv(); err != nil {
				return fmt.Errorf("failed to load environment, use 'niflheim init' to create a skeleton environment file: %s", err)
			}
			data, err := stdData()
			if err != nil {
				return fmt.Errorf("getting standard data: %w", err)
			}

			folders, err := os.ReadDir(data["VALHEIM_LOGS_FOLDER"].(string))
			if err != nil {
				return fmt.Errorf("failed to read logs folder: %s", err)
			}

			sort.Slice(folders, func(i, j int) bool {
				return folders[i].Name() > folders[j].Name()
			})
			fmt.Println("Latest log name:", folders[0].Name())

			logName := fmt.Sprintf("%s/%s", data["VALHEIM_LOGS_FOLDER"], folders[0].Name())
			t, err := tail.TailFile(logName, tail.Config{Follow: true})
			if err != nil {
				return fmt.Errorf("failed to tail log: %s", err)
			}
			for line := range t.Lines {
				fmt.Println(line.Text)
			}
			return nil

		},
	}
	app.Commands = append(app.Commands, &init)
}
