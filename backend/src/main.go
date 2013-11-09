package main

import (
    "db"
    "flag"
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
    filepath := fmt.Sprintf("%s/../frontend/src/", path)
    web.Config.StaticDir = filepath

    // Setup the DB
    db.Init()

    // Parse command line arguments
    port := flag.Int("port", 80, "port number to start the server on")
    flag.Parse()
    url := fmt.Sprintf("0.0.0.0:%d", *port)

    // Start the server!
    serve.Routes()
    web.Run(url)
}
