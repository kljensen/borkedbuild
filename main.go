package main

import (
	"./svg"
	"github.com/codegangsta/martini"
)

const defaultstyle = "fill:rgb(127,0,0)"

func main() {

	// Use martini Classic for basic defaults
	m := martini.Classic()

	// Set up our Routes
	//

	// Index
	m.Get("/", func() string {
		return "You just got Borked!"
	})

	// Borked Image Route
	m.Get("/:wordleft/:wordright/:color", borkedbuild.MakeSvg)

	// Run the server
	m.Run()
}
