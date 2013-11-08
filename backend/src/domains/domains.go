package domains

type User struct {
    PublicId int    `json:"public_id"`
    Name     string `json:"name"`
    Email    string `json:"email"`
}

type Flag struct {
    Tag        string `json:"tag"`
    Value      int    `json:"value"`
    Discovered bool   `json:"discovered"`
}

type UsersFlagInfo struct {
    User  User   `json:"user"`
    Flags []Flag `json:"flags"`
}

type Scoreboard struct {
    Scores []UsersFlagInfo `json:"scores"`
}
