package db

import (
    l4g "code.google.com/p/log4go"
    "database/sql"
    "domains"
    "fmt"
)

func GetFlagsForUser(uId int) ([]domains.Flag, error) {
    flags := make([]domains.Flag, 0)
    results, err := QueryRows("SELECT tag, value FROM Users u, Flags f, UsersFlags uf WHERE u.id = uf.uId AND f.id = uf.fId AND u.publicId = ?",
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

func GetPublicFlagsForUser(uId int) ([]domains.PublicFlag, error) {
    publicFlags := make([]domains.PublicFlag, 0)
    results, err := QueryRows("SELECT value FROM Users u, Flags f, UsersFlags uf WHERE u.id = uf.uId AND f.id = uf.fId AND u.publicId = ?",
        Params(uId), rowToPublicFlag)

    if err != nil {
        return publicFlags, err
    }

    for _, result := range results {
        flag := result.(domains.PublicFlag)
        publicFlags = append(publicFlags, flag)
    }

    return publicFlags, err
}

// Return true if flag is validated
func ValidateFlagFor(tag string, publicId int) (bool, error) {

    result, err := QueryRows("SELECT id FROM Flags WHERE tag = ?", Params(tag), rowToInt)
    if err != nil {
        return false, err
    }

    if len(result) == 0 {
        return false, nil
    }

    flagId := result[0].(int)

    // Get users DB id
    result, err = QueryRows("SELECT id FROM Users WHERE publicId = ?", Params(publicId), rowToInt)
    if err != nil {
        return false, err
    }

    if len(result) == 0 { // User not found
        return false, nil
    }

    userId := result[0].(int)

    // Check if it was already validated
    result, err = QueryRows("SELECT uId FROM UsersFlags WHERE fId = ?", Params(flagId), rowToInt)
    if err != nil {
        return false, err
    }

    if len(result) != 0 { // Flag found
        return false, nil
    }

    // Set the flag as found
    _, err = Insert("INSERT UsersFlags (uId, fId) VALUES (?, ?) ", Params(userId, flagId))
    if err != nil {
        errMsg := fmt.Sprintf("Failed INSERT to UsersFlags for USER %d and FLAG %d", userId, flagId)
        l4g.Error(errMsg)
        return false, err
    }

    return true, nil
}

func rowToFlag(rows *sql.Rows) (interface{}, error) {
    flag := domains.Flag{}
    err := rows.Scan(&flag.Tag, &flag.Value)
    return flag, err
}

func rowToPublicFlag(rows *sql.Rows) (interface{}, error) {
    flag := domains.PublicFlag{}
    err := rows.Scan(&flag.Value)
    return flag, err
}

func rowToInt(rows *sql.Rows) (interface{}, error) {
    var r int
    err := rows.Scan(&r)
    return r, err
}
