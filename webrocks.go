package main

import (
    "github.com/jamierocks/webrocks/controllers"
    "github.com/jamierocks/webrocks/modules"
    "github.com/go-macaron/pongo2"
    "gopkg.in/macaron.v1"
)

func main() {
    // Initialise config
    modules.InitConfig()

    // Initialise database
    modules.InitDatabase()

    // Initialise Macaron
    m := macaron.Classic()
    m.Use(pongo2.Pongoer())

    // Register routes
    m.Get("/", controllers.GetIndex)
    m.NotFound(controllers.NotFound)

    // Start Macaron
    m.Run()
}
