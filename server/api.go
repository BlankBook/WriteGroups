package server

import (
    "net/http"
    "database/sql"

    "github.com/blankbook/shared/models"
    "github.com/blankbook/shared/web"
)

// SetupAPI adds the API routes to the provided router
func SetupAPI(r web.Router, db *sql.DB) {
    r.HandleRoute([]string{web.POST}, "/group",
                  []string{},
                  []string{},
                  PostGroup, db)
}

func PostGroup(w http.ResponseWriter, q map[string][]string, b string, db *sql.DB) {
    var err error
    defer func() {
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
        }
    }()
    g, err := models.ParseGroup(b)
    if err != nil {
        return
    }
    err = g.Validate()
    if err != nil {
        return
    }
    query := `INSERT INTO Groups (Name, Protected) Values ($1, $2)`
    protectedVal := 0
    if g.Protected {
        protectedVal = 1
    }
    _, err = db.Exec(query, g.Name, protectedVal)
    if err != nil {
        return
    }
    w.WriteHeader(http.StatusOK)
}
