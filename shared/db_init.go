package shared

import (
    "database/sql"
    "golang-echo-showcase/src/user/sqlc/output"
    _ "github.com/mattn/go-sqlite3" 
)

type Database struct {
    Queries *sqlc.Queries
    DB      *sql.DB
}

func NewDatabase(dataSourceName string) (*Database, error) {
    database, err := sql.Open("sqlite3", dataSourceName)
    if err != nil { return nil, err }
    
    queries := sqlc.New(database)
    return &Database{Queries: queries, DB: database}, nil
}

func (d* Database) CloseDatabase() error {
    if d.DB != nil {
        return d.DB.Close()
    }
    return nil
}
