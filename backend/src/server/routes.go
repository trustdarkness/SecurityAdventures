package serve

import (
    "github.com/hoisie/web"
)

func Routes() {

    // "Brain" stuff
    web.Get("/scoreboard", getScoreboard)
    web.Get("/users", getUsers)
    web.Put("/validateFlag", domainHandler(validateFlag))

    // To be implemented ?
    web.Post("/newuser", domainHandler(newUser))
}
