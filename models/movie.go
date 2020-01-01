package models

import (
	"github.com/globalsign/mgo/bson"
)

// Movie struct
type Movie struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	Description string        `bson:"description" json:"description"`
}

const (
	db         = "movie"
	collection = "movies"
)

// FindAll will find all movies.
func (m *Movie) FindAll() ([]Movie, error) {
	var movies []Movie
	err := FindAll(db, collection, nil, nil, &movies)

	return movies, err
}

// FindByID will find a movie by ID.
func (m *Movie) FindByID(id string) (Movie, error) {
	var movie Movie
	err := FindByID(db, collection, id, &movie)

	return movie, err
}

// Store will store a movie.
func (m *Movie) Store(movie Movie) error {
	return Insert(db, collection, movie)
}

// Update will update a movie.
func (m *Movie) Update(id string, movie Movie) error {
	return UpdateByID(db, collection, id, movie)
}

// Remove will remove a movie.
func (m *Movie) Remove(id string) error {
	return RemoveByID(db, collection, id)
}
