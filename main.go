package main

import (
	"fmt"
	"os"

	"github.com/pharrisee/niflheim/actions"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Usage: "setup a dedicated Valheim server on linux",
	}

	actions.Init(app)
	actions.Depends(app)
	actions.Install(app)
	actions.Env(app)
	actions.ServiceInstall(app)
	actions.ServiceStart(app)
	actions.ServiceStatus(app)
	actions.ServiceStop(app)
	actions.ServiceRestart(app)
	actions.Tail(app)

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("Error encountered: %v\n", err)
	}
}
