package main

import (
    "gopkg.in/macaron.v1"
)

func main() {
    // Initialise Macaron
    m := macaron.Classic()

    // Start Macaron
    m.Run()
}
