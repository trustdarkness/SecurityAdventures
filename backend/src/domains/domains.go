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
    Tag   string `json:"tag"`
    Value int    `json:"value"`
}

type PublicFlag struct {
    Value int `json:"value"`
}

type ValidateFlag struct {
    PublicUserId int    `json:"public_user_id"`
    Tag          string `json:"tag"`
}

type UsersFlagInfo struct {
    User  PublicUser `json:"user"`
    Flags []Flag     `json:"flags"`
}

type UsersPublicFlagInfo struct {
    PublicUser  PublicUser   `json:"public_user"`
    PublicFlags []PublicFlag `json:"public_flags"`
}

type Scoreboard struct {
    Scores []UsersPublicFlagInfo `json:"scores"`
}

func BytesToFlag(b []byte) (Flag, error) {
    flag := Flag{}
    err := json.Unmarshal(b, &flag)
    return flag, err
}

func BytesToPublicFlag(b []byte) (PublicFlag, error) {
    publicFlag := PublicFlag{}
    err := json.Unmarshal(b, &publicFlag)
    return publicFlag, err
}

func BytesToValidateFlag(b []byte) (ValidateFlag, error) {
    validateFlag := ValidateFlag{}
    err := json.Unmarshal(b, &validateFlag)
    return validateFlag, err
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
    userFlagInfo := UsersFlagInfo{}
    err := json.Unmarshal(b, &userFlagInfo)
    return userFlagInfo, err
}
