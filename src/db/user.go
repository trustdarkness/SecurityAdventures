package db

import (
    "database/sql"
    "domains"
)

// External functions
func GetUser(uId int) (domains.User, error) {
    result, err := QueryRows("SELECT publicId, name, email FROM Users WHERE publicId = ?",
        Params(uId), rowToUser)

    if len(result) == 0 || err != nil {
        return domains.User{}, err
    }

    user := result[0].(domains.User)
    return user, err
}

func GetUsers() ([]domains.User, error) {
    users := make([]domains.User, 0)
    results, err := QueryRows("SELECT publicId, name, email FROM Users",
        nil, rowToUser)

    if err != nil {
        return users, err
    }

    for _, result := range results {
        user := result.(domains.User)
        users = append(users, user)
    }

    return users, nil
}

// Transform funcs
func rowToUser(rows *sql.Rows) (interface{}, error) {
    user := domains.User{}
    err := rows.Scan(&user.PublicId, &user.Name, &user.Email)
    return user, err
}
