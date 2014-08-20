package dbutil

import (
    "database/sql"
)

func Insert(statement string, args ...interface{}) (uint64, error) {
    dbconn := Get()
    id := uint64(0)
    err := dbconn.QueryRow(statement, args...).Scan(&id)
    return id, err
}

func Select(f ParseRowFunc, statement string, args ...interface{}) (Object, error) {
    var row *sql.Row
    dbconn := Get()
    if len(args) > 0 {
        row = dbconn.QueryRow(statement, args...)
    } else {
        row = dbconn.QueryRow(statement)
    }
    return f(row)
}

func SelectList(f ParseRowFunc, statement string, args ...interface{}) ([]Object, error) {
    var err error
    var rows *sql.Rows
    var item Object
    var itemList []Object
    dbconn := Get()
    if len(args) > 0 {
        rows, err = dbconn.Query(statement, args)
    } else {
        rows, err = dbconn.Query(statement)
    }
    if err != nil { return itemList, err }
    defer rows.Close()
    for rows.Next() {
        item, err = f(rows)
        if err != nil { break }
        itemList = append(itemList, item)
    }
    return itemList, rows.Err()
}

