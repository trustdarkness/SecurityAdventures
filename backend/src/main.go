package main

import (
    "db"
    "fmt"
    "github.com/hoisie/web"
    "os"
    "server"
)

func main() {

    // Set the static directory for webgo
    path := os.Getenv("GOPATH")
    if path == "" {
        fmt.Println("GOPATH NOT SET")
        return
    }
    filepath := fmt.Sprintf("%s/../frontend/", path)
    web.Config.StaticDir = filepath

    // Setup the DB
    db.Init()

    // Start the server!
    serve.Routes()
    web.Run("0.0.0.0:9999")
}
