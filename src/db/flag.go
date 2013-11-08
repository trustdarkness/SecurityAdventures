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

func PublishFlag(flag domains.Flag) error {
    return nil
}

func rowToFlag(rows *sql.Rows) (interface{}, error) {
    flag := domains.Flag{}
    err := rows.Scan(&flag.Tag, &flag.Value, &flag.Discovered)
    return flag, err
}
