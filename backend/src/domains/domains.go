package domains

import (
    "encoding/json"
)

type User struct {
    PublicId int    `json:"public_id"`
    Name     string `json:"name"`
    Email    string `json:"email"`
}

type PublicUser struct {
    PublicId int `json:"public_id"`
}

type Flag struct {
    Tag        string `json:"tag"`
    Value      int    `json:"value"`
    Discovered bool   `json:"discovered"`
}

type UsersFlagInfo struct {
    User  PublicUser `json:"user"`
    Flags []Flag     `json:"flags"`
}

type Scoreboard struct {
    Scores []UsersFlagInfo `json:"scores"`
}

func BytesToFlag(b []byte) (Flag, error) {
    flag := Flag{}
    err := json.Unmarshal(b, &flag)
    return flag, err
}

func BytesToUser(b []byte) (User, error) {
    user := User{}
    err := json.Unmarshal(b, &user)
    return user, err
}

func BytesToPublicUser(b []byte) (PublicUser, error) {
    publicUser := PublicUser{}
    err := json.Unmarshal(b, &publicUser)
    return publicUser, err
}

func BytesToUsersFlagInfo(b []byte) (UsersFlagInfo, error) {
    userFlagInfp := UsersFlagInfo{}
    err := json.Unmarshal(b, &userFlagInfp)
    return userFlagInfp, err
}
