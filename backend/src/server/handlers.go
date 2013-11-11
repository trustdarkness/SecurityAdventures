package serve

import (
    "bytes"
    l4g "code.google.com/p/log4go"
    "db"
    "domains"
    "encoding/json"
    "fmt"
    "github.com/hoisie/web"
)

var errResponse = "{ \"msg\": \"have you ever danced with the devil in the pale moonlight?\" }"

type PostHandlerFunc func([]byte) string
type ExecDomainHandlerFunc func(*web.Context) string

func domainHandler(handler PostHandlerFunc) ExecDomainHandlerFunc {
    return func(ctx *web.Context) string {
        buf := new(bytes.Buffer)
        buf.ReadFrom(ctx.Request.Body)
        return handler(buf.Bytes())
    }
}

// GET REQUESTS -----------------------------------------------------------------------------------
func getScoreboard(ctx *web.Context) string {
    scoreBoard := domains.Scoreboard{}

    users, err := db.GetPublicUsers()
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

func getPublicUsers(ctx *web.Context) string {
    users, err := db.GetPublicUsers()
    return constructGetResponse("users", users, err)
}

// POST/PUT REQUESTS -----------------------------------------------------------------------------------

// These requests only request a success / failure repsonse
func validateFlag(b []byte) string {
    userFlagInfo, err := domains.BytesToUsersFlagInfo(b)
    if err != nil {
        return constructStandardResponse("", err)
    }

    found, err := db.ValidateFlagFor(userFlagInfo.Flags[0].Tag, userFlagInfo.User.PublicId)
    if err != nil {
        return constructStandardResponse("", err)
    }

    if found == true {
        return constructStandardResponse("flag validated", nil)
    }

    return constructStandardResponse("flag not validated", nil)
}

// To be implemented ?
func newUser(b []byte) string {
    return constructStandardResponse("", nil)
}

// RESPONSES CONSTRUCTION ------------------------------------------------------------------------------

func constructGetResponse(key string, out interface{}, err error) string {

    if err != nil {
        // logging errors
        l4g.Error(err)
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

func constructStandardResponse(message string, err error) string {

    if err != nil {
        // logging errors
        l4g.Error(err)
        return errResponse
    }

    return fmt.Sprintf("{ \"msg\": \"%s\" }", message)
}
