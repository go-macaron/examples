package main

import (
	"gopkg.in/macaron.v1"
)

func main() {
	m := macaron.Classic()
	m.Use(macaron.Renderer(macaron.RenderOptions{
		IndentJSON: true,
	}))

	m.Group("/songs", func() {
		m.Combo("").
			Post(CreateNewSong).
			Get(GetAllSongs)
		m.Combo("/:id").
			Get(GetSong).
			Patch(EditSong).
			Delete(DeleteSong)
	})

	m.Run()
}
