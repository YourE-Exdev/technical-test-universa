package main

import (
	"encoding/json"
	"log"
	"net/http"
	"server/models"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

type MusicPayLoad struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Album       string `json:"album"`
	AlbumArt    string `json:"albumart"`
	Singer      string `json:"singer"`
	PublishDate string `json:"publishdate"`
	CreateAt    string `json:"createat"`
	UpdateAt    string `json:"updateat"`
}

func (app *application) getOneMusic(rw http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		// app.logger.Print(errors.New("Invalid id paramater"))
		app.errorJSON(rw, err)
		return
	}

	app.logger.Println(app.models.DB.Get(id))

	music, err := app.models.DB.Get(id)

	err = app.writeJSON(rw, http.StatusOK, music, "music")
}

func (app *application) getAllMusic(rw http.ResponseWriter, r *http.Request) {

	app.logger.Println(app.models.DB.All())
	music, err := app.models.DB.All()
	if err != nil {
		app.errorJSON(rw, err)
		return
	}

	err = app.writeJSON(rw, http.StatusOK, music, "music")
	if err != nil {
		app.errorJSON(rw, err)
		return
	}
}

func (app *application) AddMusic(rw http.ResponseWriter, r *http.Request) {
	var payload MusicPayLoad
	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		log.Println(err)
		app.errorJSON(rw, err)
		return
	}

	var music models.Music

	music.ID, _ = strconv.Atoi(payload.ID)
	music.Name = payload.Name
	music.Album = payload.Album
	music.AlbumArt = payload.AlbumArt
	music.Singer = payload.Singer
	music.PublishDate = payload.PublishDate
	music.CreateAt = time.Now()
	music.UpdateAt = time.Now()

	//insert music query
	err = app.models.DB.InsertMusic(music)

	if err != nil {
		log.Println(err)
		app.errorJSON(rw, err)
		return
	}

	type jsonRes struct {
		OK bool `json:"ok"`
	}

	ok := jsonRes{
		OK: true,
	}

	err = app.writeJSON(rw, http.StatusOK, ok, "response")

	if err != nil {
		log.Println(err)
		app.errorJSON(rw, err)
		return
	}
}

func (app *application) EditMusic(rw http.ResponseWriter, r *http.Request) {
	var payload MusicPayLoad
	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		log.Println(err)
		app.errorJSON(rw, err)
		return
	}

	var music models.Music

	id, _ := strconv.Atoi(payload.ID)
	singleMusic, _ := app.models.DB.Get(id)
	music = *singleMusic

	music.ID, _ = strconv.Atoi(payload.ID)
	music.Name = payload.Name
	music.Album = payload.Album
	music.AlbumArt = payload.AlbumArt
	music.Singer = payload.Singer
	music.PublishDate = payload.PublishDate
	music.UpdateAt = time.Now()

	//update music query
	err = app.models.DB.UpdateMusic(music)

	if err != nil {
		log.Println(err)
		app.errorJSON(rw, err)
		return
	}

	type jsonRes struct {
		OK bool `json:"ok"`
	}

	ok := jsonRes{
		OK: true,
	}

	err = app.writeJSON(rw, http.StatusOK, ok, "response")

	if err != nil {
		log.Println(err)
		app.errorJSON(rw, err)
		return
	}
}

func (app *application) DeleteMusic(rw http.ResponseWriter, r *http.Request) {
	var payload MusicPayLoad
	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		log.Println(err)
		app.errorJSON(rw, err)
		return
	}

	id, _ := strconv.Atoi(payload.ID)
	err = app.models.DB.DeleteMusic(id)

	if err != nil {
		log.Println(err)
		app.errorJSON(rw, err)
		return
	}

	type jsonRes struct {
		OK bool `json:"ok"`
	}

	ok := jsonRes{
		OK: true,
	}

	err = app.writeJSON(rw, http.StatusOK, ok, "response")

	if err != nil {
		log.Println(err)
		app.errorJSON(rw, err)
		return
	}
}
