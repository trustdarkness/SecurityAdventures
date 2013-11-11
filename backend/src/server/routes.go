package serve

import (
    "github.com/hoisie/web"
)

func Routes() {

    // "Brain" stuff
    web.Get("/scoreboard", getScoreboard)
    web.Get("/publicUsers", getPublicUsers)
    web.Put("/validateFlag", domainHandler(validateFlag))

    // To be implemented ?
    web.Post("/newuser", domainHandler(newUser))
}
