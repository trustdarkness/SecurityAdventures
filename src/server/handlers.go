package serve

import (
    "db"
    "domains"
    "encoding/json"
    "fmt"
)

var errResponse = "{ \"msg\": \"have you ever danced with the devil in the pale moonlight?\" }"

func getScoreboard() string {
    scoreBoard := domains.Scoreboard{}

    users, err := db.GetUsers()
    if err != nil {
        return constructResponse("scoreboard", scoreBoard, err)
    }

    for _, user := range users {

        flags, err := db.GetFlagsForUser(user.PublicId)
        if err != nil {
            return constructResponse("scoreboard", scoreBoard, err)
        }
        flagInfo := domains.UsersFlagInfo{}
        flagInfo.User = user
        flagInfo.Flags = flags
        scoreBoard.Scores = append(scoreBoard.Scores, flagInfo)
    }

    return constructResponse("scoreboard", scoreBoard, nil)
}

func publishFlag() string {
    return constructResponse("", nil, nil)
}

func getUsers() string {
    users, err := db.GetUsers()
    return constructResponse("users", users, err)
}

// To be implemented ?
func newUser() string {
    return constructResponse("", nil, nil)
}

func constructResponse(key string, out interface{}, err error) string {
    if err != nil {
        fmt.Println(err)
        return errResponse
    }

    if out == nil {
        return errResponse
    }

    b, err := json.Marshal(out)
    if err != nil {
        return errResponse
    }
    return fmt.Sprintf("{ \"%s\": %s }", key, string(b))
}
