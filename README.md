# Lintukoto oraakkeli twitch bot
<!---[![Build Status](https://travis-ci.org/laam4/mariomaker-twitch.svg?branch=master)](https://travis-ci.org/laam4/mariomaker-twitch))-->
Code Based on https://github.com/laam4/mariomaker-twitch

## Features
- Lintukoto oraakkeli

## Install
*Warning:* Code is untested, I've never used Go and I don't know how the toolchain works. So the Code and these install instructions might not work at all :D

- Go to your go project folder
- Get dependencies `go get github.com/fatih/color` `go get github.com/vharitonsky/iniflags`
- Get twitch bot `go get github.com/fld/oraakkeli-twitch`
- Edit `default.ini`
- Type `go install github.com/fld/oraakkeli-twitch`
- Run bot from your GOPATH/bin folder with -config parameter
