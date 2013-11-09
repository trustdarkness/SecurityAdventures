package serve

import (
    "bytes"
    "db"
    "domains"
    "encoding/json"
    "fmt"
    "github.com/hoisie/web"
)

var successResponse = "{ \"msg\": \"success\" }"
var errResponse = "{ \"msg\": \"have you ever danced with the devil in the pale moonlight?\" }"

// GET REQUESTS -----------------------------------------------------------------------------------
func getScoreboard(ctx *web.Context) string {
    scoreBoard := domains.Scoreboard{}

    users, err := db.GetUsers()
    if err != nil {
        return constructGetResponse("scoreboard", scoreBoard, err)
    }

    for _, user := range users {

        flags, err := db.GetFlagsForUser(user.PublicId)
        if err != nil {
            return constructGetResponse("scoreboard", scoreBoard, err)
        }
        flagInfo := domains.UsersFlagInfo{}
        flagInfo.User = user
        flagInfo.Flags = flags
        scoreBoard.Scores = append(scoreBoard.Scores, flagInfo)
    }

    return constructGetResponse("scoreboard", scoreBoard, nil)
}

func getUsers(ctx *web.Context) string {
    users, err := db.GetUsers()
    return constructGetResponse("users", users, err)
}

// POST/PUT REQUESTS -----------------------------------------------------------------------------------
// These requests only request a success / failure repsonse
func validateFlag(ctx *web.Context) string {
    return constructStandardResponse(nil)
}

// To be implemented ?
func newUser(ctx *web.Context) string {
    return constructStandardResponse(nil)
}

// RESPONSES CONSTRUCTION ------------------------------------------------------------------------------

func constructGetResponse(key string, out interface{}, err error) string {
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

func constructStandardResponse(err error) string {

    if err != nil {
        fmt.Println(err)
        return errResponse
    }

    return successResponse
}
