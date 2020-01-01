package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
	"github.com/memochou1993/movies-api/models"
)

var (
	model = models.Movie{}
)

func response(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Index display a listing of the resource.
func Index(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var movies []models.Movie

	movies, err := model.FindAll()
	if err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	response(w, http.StatusOK, movies)
}

// Show display the specified resource.
func Show(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	movie, err := model.FindByID(id)

	if err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	response(w, http.StatusOK, movie)
}

// Store store a newly created resource in storage.
func Store(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var movie models.Movie
	movie.ID = bson.NewObjectId()

	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		response(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := movie.Store(movie); err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	response(w, http.StatusCreated, movie)
}

// Update update the specified resource in storage.
func Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	var movie models.Movie
	movie.ID = bson.ObjectIdHex(id)

	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		response(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := model.Update(id, movie); err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	response(w, http.StatusOK, movie)
}

// Destroy remove the specified resource from storage.
func Destroy(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	if err := model.Remove(id); err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	response(w, http.StatusNoContent, nil)
}
