package server

import (
    "os"
    "net/http"
    "log"

    "github.com/gorilla/mux"

    "github.com/blankbook/shared/web"
)

const pathPrefix = "/groups/write"
const databaseUsernameEnvVar = "BB_GROUPS_DB_USERNAME"
const databasePasswordEnvVar = "BB_GROUPS_DB_PASSWORD"
const databaseServerEnvVar = "BB_GROUPS_DB_SERVER"
const dbName = "blankbookgroups"

// SetupRoutes configures the service API endpoints
func SetupRoutes() {
    dbUsername := os.Getenv(databaseUsernameEnvVar)
    dbPassword := os.Getenv(databasePasswordEnvVar)
    dbServer := os.Getenv(databaseServerEnvVar)
    db, err := web.GetMSSqlDatabase(dbUsername, dbPassword, dbServer, dbName)
    if err != nil {
        log.Panic(err.Error())
    }

    muxRouter := mux.NewRouter()
    muxRouter.NotFoundHandler = http.HandlerFunc(notFoundHandler)
    r := web.NewHTTPRouter(muxRouter, pathPrefix)
    SetupAPI(r, db)
    r.StartListening()
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "endpoint not found", http.StatusNotFound)
}
