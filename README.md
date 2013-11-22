#CTF Server/Scoreboard Project
####Chris Dargis
[Security Adventures](http://securityadventures.org/)

##Dependencies
+ [NPM](https://npmjs.org/)
+ [Coffeescript](http://coffeescript.org/)
+ [LESS CSS](http://lesscss.org/)
+ [Go](http://golang.org/)
  * [Go-MySQL-Driver](http://godoc.org/github.com/go-sql-driver/mysql)
  * [web.go](http://webgo.io/)
  * [log4go](http://code.google.com/p/log4go)
+ MySQL

##Setup
+ Frontend
  * Install NPM
  * Install coffeescript: `npm install -g coffee-script`
  * Install LESS CSS: `npm install -g less`
  * Run `sh frontend/compileFrontend.sh`
+ Backend
  * Install Go
  * Set GOPATH to `PATH-TO-THIS-REPO/backend`
  * Run `sh backend/getGoDeps.sh`
  * Setup MySQL
    - `CREATE DATABASE SecurityAdventures;`
    - `CREATE USER 'USER_NAME'@'localhost' IDENTIFIED BY 'PASSWORD';`
    - `GRANT ALL PRIVILEGES ON SecurityAdventures . * TO 'USER_NAME'@localhost';`
    - Update `db.go` with USER_NAME and PASSWORD
    - Initialize the Schema: `mysql --user=USER_NAME --password=PASSWORD SecurityAdventures < SecurityAdventures/backend/schema.sql`
    - Use the shell script to generate SQL: `sh SecurityAdventures/backend/generateSQL.sh`
    - Run generated SQL script: `mysql --user=USER_NAME --password=PASSWORD SecurityAdventures < seed.sql`
    - Create a `dbconfig.txt` file that looks something like this:

```
User = "AwesomeUser"
Pass = "AwesomeUsersPassword"
DBName = "SecurityAdventures"
```

##Building the Backend
Simple: `go build -o scoreboard SecurityAdventures/backend/src/main.go`

##Running the Server
You can run with `go run` command:

from `SecurityAdventures/` run `go run backend/src/main.go`

Or simply launch the build:

`./scoreboard > traffic_log.txt`

You can also pass the port numbeer for the server to listen on:

`./scoreboard --port=80`
