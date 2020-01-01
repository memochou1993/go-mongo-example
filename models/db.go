package models

import (
	"log"

	"github.com/globalsign/mgo"
)

const (
	host     = "localhost:27017"
	source   = "movies"
	username = ""
	password = ""
)

var globalS *mgo.Session

func init() {
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{host},
		Source:   source,
		Username: username,
		Password: password,
	}

	s, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	globalS = s
}

func connect(db string, collection string) (*mgo.Session, *mgo.Collection) {
	s := globalS.Copy()
	c := s.DB(db).C(collection)

	return s, c
}

// FindAll will find all resources.
func FindAll(db string, collection string, query, selector, result interface{}) error {
	s, c := connect(db, collection)
	defer s.Close()

	return c.Find(query).Select(selector).All(result)
}

// FindOne will find a resource.
func FindOne(db string, collection string, query, selector, result interface{}) error {
	s, c := connect(db, collection)
	defer s.Close()

	return c.Find(query).Select(selector).One(result)
}

// Insert will insert a resource.
func Insert(db string, collection string, docs ...interface{}) error {
	s, c := connect(db, collection)
	defer s.Close()

	return c.Insert(docs...)
}

// Update will update a resource.
func Update(db string, collection string, query, update interface{}) error {
	s, c := connect(db, collection)
	defer s.Close()

	return c.Update(query, update)
}

// UpdateByID will update a resource By ID.
func UpdateByID(db string, collection string, id interface{}, update interface{}) error {
	s, c := connect(db, collection)
	defer s.Close()

	return c.UpdateId(id, update)
}

// Remove will remove a resource.
func Remove(db string, collection string, query interface{}) error {
	s, c := connect(db, collection)
	defer s.Close()

	return c.Remove(query)
}

// RemoveByID will remove a resource by ID.
func RemoveByID(db string, collection string, id interface{}) error {
	s, c := connect(db, collection)
	defer s.Close()

	return c.RemoveId(id)
}
