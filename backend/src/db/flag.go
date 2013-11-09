package db

import (
    "database/sql"
    "domains"
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
func ValidateFlag(tag string) (bool, error) {
    result, err := QueryRows("SELECT id FROM Flags f WHERE discovered = 0 AND tag = ?", Params(tag), rowToInt)

    if err != nil {
        return false, err
    }

    if len(result) == 0 {
        return false, nil
    }

    id := result[0].(int)
    err = UpdateRow("UPDATE Flags SET discovered = 1 WHERE id = ?", Params(id))

    if err != nil {
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
