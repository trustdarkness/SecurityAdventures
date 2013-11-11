package db

import (
    l4g "code.google.com/p/log4go"
    "database/sql"
    "domains"
    "fmt"
)

func GetFlagsForUser(uId int) ([]domains.Flag, error) {
    flags := make([]domains.Flag, 0)
    results, err := QueryRows("SELECT tag, value, discovered FROM Users u, Flags f, UsersFlags uf WHERE u.id = uf.uId AND f.id = uf.fId AND u.publicId = ?",
        Params(uId), rowToFlag)

    if err != nil {
        return flags, err
    }

    for _, result := range results {
        flag := result.(domains.Flag)
        flags = append(flags, flag)
    }

    return flags, err
}

// Return true if flag is validated
func ValidateFlagFor(tag string, publicId int) (bool, error) {
    result, err := QueryRows("SELECT id FROM Flags WHERE discovered = 0 AND tag = ?", Params(tag), rowToInt)

    if err != nil {
        return false, err
    }

    if len(result) == 0 {
        return false, nil
    }

    flagId := result[0].(int)

    result, err = QueryRows("SELECT id FROM Users WHERE publicId = ?", Params(publicId), rowToInt)
    if len(result) == 0 {
        return false, nil
    }

    userId := result[0].(int)

    _, err = Insert("INSERT UsersFlags VALUES (?, ?) ", Params(userId, flagId))
    if err != nil {
        errMsg := fmt.Sprintf("Failed INSERT to UsersFlags for USER %d and FLAG %d", userId, flagId)
        l4g.Error(errMsg)
        return false, err
    }

    err = UpdateRow("UPDATE Flags SET discovered = 1 WHERE id = ?", Params(flagId))

    if err != nil {
        errMsg := fmt.Sprintf("Failed UPDATE to Flags for USER %d and FLAG %d - trying to SET flag discovered",
            userId, flagId)
        l4g.Error(errMsg)
        return false, err
    }

    return true, nil
}

func rowToFlag(rows *sql.Rows) (interface{}, error) {
    flag := domains.Flag{}
    err := rows.Scan(&flag.Tag, &flag.Value, &flag.Discovered)
    return flag, err
}

func rowToInt(rows *sql.Rows) (interface{}, error) {
    var r int
    err := rows.Scan(&r)
    return r, err
}
