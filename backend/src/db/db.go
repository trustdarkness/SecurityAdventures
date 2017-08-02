package db

import (
    "bufio"
    l4g "github.com/alecthomas/log4go"
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "os"
    "strings"
)

type DBConfig struct {
    User   string
    Pass   string
    DBName string
}

type RowTransformFunc func(*sql.Rows) (interface{}, error)

var dbConfig DBConfig

func Init(path string) {
    fullPath := fmt.Sprintf("%s/dbconfig.txt", path)
    file, err := os.Open(fullPath)
    if err != nil {
        l4g.Error("Could not open dbconfig.txt")
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        tokens := strings.Split(scanner.Text(), "=")
        for index, token := range tokens {
            token = strings.Trim(token, " \"")
            tokens[index] = token
        }
        if len(tokens) == 2 {
            key := tokens[0]
            value := tokens[1]
            if key == "User" {
                dbConfig.User = value
            }
            if key == "Pass" {
                dbConfig.Pass = value
            }
            if key == "DBName" {
                dbConfig.DBName = value
            }
        }
    }
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
