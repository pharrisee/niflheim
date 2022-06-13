# niflheim

Easy setup of a Valheim server (currently Ubuntu 20.04 specific).

NAME:
   niflheim - setup a dedicated Valheim server on linux

USAGE:
   niflheim [global options] command [command options] [arguments...]

COMMANDS:
   init             create a niflheim environment file
   depends          install required dependencies
   install          install a Valheim instance
   env              print the current environment
   service-install  install the service for the instance
   service-start    start the service for the instance
   service-status   status of the service for the instance
   service-stop     stop the service for the instance
   service-restart  start the service for the instance
   tail             tail the log
   help, h          Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)

## Plan of attack

1. <a href="https://github.com/pharrisee/niflheim/releases">Download a release</a>
2. move it to a folder of you chooising, preferably on the path
3. run it, you'll get some help text

## init

the `niflheim init` command will initialise a `niflheim.env` file in the root of your current users folder.  Take a read of it, and edit it to suit your environment.

```
VALHEIM_SERVER_FOLDER=/home/username/valheim-server
VALHEIM_DATA_FOLDER=/home/username/valheim-data
VALHEIM_LOGS_FOLDER=/home/username/valheim-logs
VALHEIM_START_SCRIPT_NAME=niflheim-start.sh

VALHEIM_SERVER_NAME="Niflheim"
VALHEIM_SERVER_PASSWORD="some random password"
VALHEIM_SERVER_PORT=2456
VALHEIM_SERVER_WORLD="niflheim"
VALHEIM_SERVER_PUBLIC=1

VALHEIM_SERVICE_NAME=valheim-username
```