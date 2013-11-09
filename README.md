#Security Adventures CTF Server/Scoreboard Project
####Chris Dargis

##Dependencies
+ [Coffeescript](http://coffeescript.org/)
+ [LESS CSS](http://lesscss.org/)
+ [Go](http://golang.org/)
  * [Go-MySQL-Driver](http://godoc.org/github.com/go-sql-driver/mysql)
  * [web.go](http://webgo.io/)
  * [log4go](http://code.google.com/p/log4go)

##Setup
+ Frontend
  * Install coffeescript
  * Install LESS CSS
  * Run `sh frontend/compileFrontend.sh`
+ Backend
  * Install Go
  * Set GOPATH to `PATH-TO-THIS-REPO/backend`
  * Run `sh backend/getGoDeps.sh`

##Building the Backend
Simple: `go build -o scoreboard SecurityAdventures/backend/src/main.go`

##Running the Server
You can run with `go run` command:

from `SecurityAdventures/` run `go run backend/src/main.go`

Or simply launch the build:

`./scoreboard > traffic_log.txt`

You can also pass the port numbeer for the server to listen on:

`./scoreboard --port=80`