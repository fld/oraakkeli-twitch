# Super Mario Maker twitch bot
<!---[![Build Status](https://travis-ci.org/laam4/mariomaker-twitch.svg?branch=master)](https://travis-ci.org/laam4/mariomaker-twitch))-->

Twitch bot written in golang which collects Super Mario Maker level codes and adds them to MySQL database.

IRC code Based on https://github.com/Vaultpls/Twitch-IRC-Bot

## Features
- Lintukoto oraakkeli

## Install
- Go to your go project folder
- Get dependencies `go get github.com/fatih/color` `go get github.com/vharitonsky/iniflags`
- Get twitch bot `go get github.com/fld/oraakkeli-twitch`
- Edit `default.ini`
- Type `go install github.com/fld/oraakkeli-twitch`
- Run bot from your GOPATH/bin folder with -config parameter
