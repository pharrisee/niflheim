package actions

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func Install(app *cli.App) {
	action := cli.Command{
		Name:  "install",
		Usage: "install a Valheim instance",
		Action: func(*cli.Context) error {
			if err := LoadEnv(); err != nil {
				return fmt.Errorf("failed to load environment, use 'niflheim init' to create a skeleton environment file: %s", err)
			}
			data, err := stdData()
			if err != nil {
				return fmt.Errorf("getting standard data: %w", err)
			}

			valplusURL, valplusName := getLatestGithubRelease("valheimplus", "valheimplus")
			cmd := fmt.Sprintf("wget -O /tmp/%s %s", valplusName, valplusURL)
			if err := runCmd(cmd); err != nil {
				return fmt.Errorf("failed to download %s: %w", valplusURL, err)
			}

			cmd = fmt.Sprintf(`/bin/sh -c 'mkdir -p %s %s'`, data["VALHEIM_DATA_FOLDER"], data["VALHEIM_LOGS_FOLDER"])
			if err := runCmd(cmd); err != nil {
				return fmt.Errorf("creating UserDataFolder: %w", err)
			}

			cmd = fmt.Sprintf(`/usr/games/steamcmd +force_install_dir %s +login anonymous +app_update 896660 +quit`, data["VALHEIM_SERVER_FOLDER"])
			if err := runCmd(cmd); err != nil {
				return fmt.Errorf("installing instance: %w", err)
			}

			cmd = fmt.Sprintf(`tar xf /tmp/%s -C %s`, valplusName, data["VALHEIM_SERVER_FOLDER"])
			if err := runCmd(cmd); err != nil {
				return fmt.Errorf("extracting instance: %w", err)
			}

			out, err := render(startScriptTemplate, data)
			if err != nil {
				return fmt.Errorf("rendering startScriptTemplate: %w", err)
			}
			scriptName := fmt.Sprintf("%s/%s", data["VALHEIM_SERVER_FOLDER"], data["VALHEIM_START_SCRIPT_NAME"])
			err = os.WriteFile(scriptName, []byte(out), 0755)
			if err != nil {
				return fmt.Errorf("writing startScriptTemplate: %w", err)
			}
			return nil
		},
	}
	app.Commands = append(app.Commands, &action)
}

var startScriptTemplate = `#!/bin/sh
log_date=$(date +'%Y%m%d%H%M')

if [ ! -f  ${HOME}/niflheim.env ]; then
	echo "niflheim.env not found, please run 'niflheim init' first"
	exit 1
fi

export $(cat $HOME/niflheim.env | xargs)

cd ${VALHEIM_SERVER_FOLDER} #  cd to the server install folder, everything after this is relative

# Whether or not to enable Doorstop. Valid values: TRUE or FALSE
export DOORSTOP_ENABLE=TRUE

# What .NET assembly to execute. Valid value is a path to a .NET DLL that mono can execute.
export DOORSTOP_INVOKE_DLL_PATH="./BepInEx/core/BepInEx.Preloader.dll"

# Which folder should be put in front of the Unity dll loading path
export DOORSTOP_CORLIB_OVERRIDE_PATH="./unstripped_corlib"

# set up LD path
# save old LD path
export templdpath="$LD_LIBRARY_PATH"

export LD_LIBRARY_PATH="./doorstop_libs":"${LD_LIBRARY_PATH}"
export LD_PRELOAD="libdoorstop_x64.so":"${LD_PRELOAD}"
export DYLD_LIBRARY_PATH="./doorstop_libs"
export DYLD_INSERT_LIBRARIES="./doorstop_libs/libdoorstop_x64.so"
export LD_LIBRARY_PATH="${VALHEIM_SERVER_FOLDER}/linux64":"${LD_LIBRARY_PATH}"

"./valheim_server.x86_64" \
	 -name "${VALHEIM_SERVER_NAME}" \
	 -password "${VALHEIM_SERVER_PASSWORD}" \
	 -port "${VALHEIM_SERVER_PORT}" \
     -world "${VALHEIM_SERVER_WORLD}" \
	 -public "${VALHEIM_SERVER_PUBLIC}" \
	 -savedir "${VALHEIM_DATA_FOLDER}" > ${VALHEIM_LOGS_FOLDER}/server-${log_date}.log 
	 -console

# restore old LD path
export LD_LIBRARY_PATH=$templdpath
`
