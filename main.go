package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func main() {
	m := martini.Classic()

	// Set a layout.
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))

	m.Get("/", func(r render.Render) {
		r.HTML(200, "hello", "jeremy")
	})

	m.Get("/main", func(r render.Render) {
		r.HTML(200, "hello", "milo")
	})

	m.Run()
}
