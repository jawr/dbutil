package dbutil

import (
    "fmt"
    _ "github.com/lib/pq"
    "github.com/jmoiron/sqlx"
)

type connection struct {
    username string
    password string
    dbname string
    db *sqlx.DB
}

var instance *connection = nil
func Setup(username, password, dbname string) error {
    if instance == nil {
        dbconn, err := sqlx.Open("postgres",
            fmt.Sprintf("user=%s password=%s dbname=%s",
                username,
                password,
                dbname,
            ),
        )
        if err != nil { return err }
        instance = &connection{
            username,
            password,
            dbname,
            dbconn,
        }
    }
    // add ability to update/reload self
    return nil
}

func Get() *sqlx.DB {
    if instance == nil {
        panic("dbutil connection instance not setup.")
    }
    if err := instance.db.Ping(); err != nil {
        // try reload
        panic(err)
    }
    return instance.db
}
