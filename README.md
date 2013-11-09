#Security Adventures CTF Server/Scoreboard Project
###Author: Chris Dargis

##Dependencies:
+ [Coffeescript](http://coffeescript.org/)
+ [LESS CSS](http://lesscss.org/)
+ [Go](http://golang.org/)
  * [Go-MySQL-Driver](http://godoc.org/github.com/go-sql-driver/mysql)
  * [web.go](http://webgo.io/)
  * [log4go](http://code.google.com/p/log4go)

##Setup:
+ Frontend
  * Install coffeescript
  * Install LESS CSS
  * Run `sh frontend/compileFrontend.sh`
+ Backend
  * Install Go
  * Set GOPATH to `PATH-TO-THIS-REPO/backend`
  * Run `sh backend/getGoDeps.sh`

##Running the Server:
From `SecurityAdventures/` run `go run backend/src/main.go`