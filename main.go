package main

import (
    "flag"
    "fmt"
    "net/http"

    "github.com/blankbook/writegroups/server"
)

func main() {
    server.SetupRoutes()

    var port int
    flag.IntVar(&port, "port", 8080, "The port to listen on")
    flag.Parse()
    http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
