package db

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

type DBConfig struct {
    User   string
    Pass   string
    DBName string
}

type RowTransformFunc func(*sql.Rows) (interface{}, error)

var dbConfig DBConfig

func Init() {
    dbConfig.User = "zeroCool"
    dbConfig.Pass = "crash"
    dbConfig.DBName = "SecurityAdventures"
}

func openDB() (*sql.DB, error) {
    openString := fmt.Sprintf("%s:%s@/%s", dbConfig.User, dbConfig.Pass, dbConfig.DBName)
    return sql.Open("mysql", openString)
}

func Insert(query string, input []interface{}) (int, error) {

    db, err := openDB()
    if err != nil {
        return -1, err
    }
    defer db.Close()

    stmt, err := db.Prepare(query)
    if err != nil {
        return -1, err
    }
    defer stmt.Close()

    res, err := stmt.Exec(input...)
    if err != nil {
        return -1, err
    }

    id, err := res.LastInsertId()

    return int(id), err
}

func QueryRows(query string, input []interface{}, rowTransform RowTransformFunc) ([]interface{}, error) {
    var results []interface{}

    db, err := openDB()
    if err != nil {
        return results, err
    }
    defer db.Close()

    stmt, err := db.Prepare(query)
    if err != nil {
        return results, err
    }
    defer stmt.Close()

    rows, err := stmt.Query(input...)
    if err != nil {
        return results, err
    }
    for rows.Next() {

        row, err := rowTransform(rows)
        if err != nil {
            return results, err
        }
        results = append(results, row)
    }
    return results, nil
}

func UpdateRow(query string, params []interface{}) error {

    db, err := openDB()
    if err != nil {
        return err
    }
    defer db.Close()

    stmt, err := db.Prepare(query)
    if err != nil {
        return err
    }
    defer stmt.Close()

    _, err = stmt.Exec(params...)
    if err != nil {
        return err
    }

    return nil
}

func Params(value ...interface{}) []interface{} {
    return value
}
