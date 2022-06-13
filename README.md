# niflheim

Easy setup of a Valheim server (currently Ubuntu 20.04 specific).

<h2>Caveat Emptor</h2>
<h3 style="color:red">Not currently for use in a production environment</h3>

```
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
```

# What it does

Currently `niflheim` will install dependencies, game server and service management of a Valheim server along with ValheimPlus and BepInEx to allow for mods to be later loaded.

`niflheim` is very early software for my own consumption, it's set up to do things as I do them, but automated.



## Plan of attack

1. <a href="https://github.com/pharrisee/niflheim/releases">Download a release</a>
2. move it to a folder of your choosing, preferably on the path.
3. run `niflheim`, you'll get some help text to peruse.

## init

The `niflheim init` command will initialise a `niflheim.env` file in the root of your current users folder.  Take a read of it, and edit it to suit your environment.

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

## depends

The `niflheim depends` command will install all of the operating system dependencies to get Valheim running quickly:

```
sudo apt-get update
sudo apt-get install --no-install-recommends --no-install-suggests -y software-properties-common
sudo dpkg --add-architecture i386
sudo add-apt-repository multiverse
echo PURGE | sudo debconf-communicate steam
echo PURGE | sudo debconf-communicate steamcmd
echo steam steam/question select "I AGREE" | sudo debconf-set-selections
echo steam steam/license note '' | sudo debconf-set-selections
echo steam steam/purge note '' | sudo debconf-set-selections
sudo apt-get install --no-install-recommends --no-install-suggests -y steamcmd lib32gcc-s1 lib32stdc++6 libsdl2-2.0-0:i386 libsdl2-2.0-0
/usr/games/steamcmd +quit &> /dev/null
sudo apt-get full-upgrade -y --allow-downgrades
sudo apt-get autoremove -y
sudo apt-get autoclean -y
```

## install

The `niflheim install` command will install valheim and also create data and logs folders for use during execution of the server.

## env

The `niflheim env` command will show you the complete environment that is used by `niflheim` when running its commands.

## service-install

`niflheim service-install` will install a systemd unitfile into /etc/systemd/system.  The actual unitfile will be named as in the `niflheim.env` created earlier.  It will also enable the service so it will restart on system start, and allow using normal service tasks such as `sudo service valheim-username stop`.
```
VALHEIM_SERVICE_NAME=valheim-username
```

## service-start

`niflheim service-start` will use the underlying `systemctl` and `service` commands to start the service.

## service-stop

`niflheim service-stop` will use the underlying `systemctl` and `service` commands to stop the service.

## service-restart

`niflheim service-restart` will use the underlying `systemctl` and `service` commands to restart the service.

## service-status

`niflheim service-status` will use the underlying `systemctl` and `service` commands to show the status of the service.

## tail

`niflheim tail` will tail the latest log from the logs folder.









