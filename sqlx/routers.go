package main

import (
	"gopkg.in/macaron.v1"
)

func CreateNewSong(ctx *macaron.Context) {
	song := &Song{
		Title:  ctx.QueryTrim("title"),
		Artist: ctx.QueryTrim("artist"),
		Year:   ctx.QueryInt("year"),
	}

	if _, err := db.Exec("INSERT INTO song (title, artist, year) VALUES ($1, $2, $3)",
		song.Title, song.Artist, song.Year); err != nil {
		ctx.JSON(500, map[string]string{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(201, song)
}

func GetAllSongs(ctx *macaron.Context) {
	songs := make([]*Song, 0, 10)
	if err := db.Select(&songs, "SELECT * FROM song"); err != nil {
		ctx.JSON(500, map[string]string{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, songs)
}

func GetSong(ctx *macaron.Context) {
	song := new(Song)
	if err := db.Get(song, "SELECT * FROM song WHERE id=$1", ctx.ParamsInt64(":id")); err != nil {
		ctx.JSON(500, map[string]string{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, song)
}

func EditSong(ctx *macaron.Context) {
	song := &Song{
		Title:  ctx.QueryTrim("title"),
		Artist: ctx.QueryTrim("artist"),
		Year:   ctx.QueryInt("year"),
	}

	if _, err := db.Exec("UPDATE song SET title=$1, artist=$2, year=$3 WHERE id=$4",
		song.Title, song.Artist, song.Year, ctx.ParamsInt64(":id")); err != nil {
		ctx.JSON(500, map[string]string{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, song)
}

func DeleteSong(ctx *macaron.Context) {
	if _, err := db.Exec("DELETE FROM song WHERE id=$1", ctx.ParamsInt64(":id")); err != nil {
		ctx.JSON(500, map[string]string{
			"error": err.Error(),
		})
		return
	}

	ctx.Status(200)
}
