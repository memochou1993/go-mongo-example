package models

import (
	"log"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

const (
	host     = "localhost:27017"
	source   = ""
	username = ""
	password = ""
)

var session *mgo.Session

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

	session = s
}

func connect(db string, collection string) (*mgo.Session, *mgo.Collection) {
	s := session.Copy()
	c := s.DB(db).C(collection)

	return s, c
}

// FindAll will find all resources.
func FindAll(db string, collection string, query interface{}, selector interface{}, result interface{}) error {
	s, c := connect(db, collection)
	defer s.Close()

	return c.Find(query).Select(selector).All(result)
}

// Find will find a resource.
func Find(db string, collection string, query interface{}, selector interface{}, result interface{}) error {
	s, c := connect(db, collection)
	defer s.Close()

	return c.Find(query).Select(selector).One(result)
}

// FindByID will find a resource by ID.
func FindByID(db string, collection string, id string, result interface{}) error {
	s, c := connect(db, collection)
	defer s.Close()

	return c.FindId(bson.ObjectIdHex(id)).One(result)
}

// Insert will insert a resource.
func Insert(db string, collection string, docs ...interface{}) error {
	s, c := connect(db, collection)
	defer s.Close()

	return c.Insert(docs...)
}

// Update will update a resource.
func Update(db string, collection string, selector interface{}, update interface{}) error {
	s, c := connect(db, collection)
	defer s.Close()

	return c.Update(selector, update)
}

// UpdateByID will update a resource By ID.
func UpdateByID(db string, collection string, id string, update interface{}) error {
	s, c := connect(db, collection)
	defer s.Close()

	return c.UpdateId(bson.ObjectIdHex(id), update)
}

// Remove will remove a resource.
func Remove(db string, collection string, selector interface{}) error {
	s, c := connect(db, collection)
	defer s.Close()

	return c.Remove(selector)
}

// RemoveByID will remove a resource by ID.
func RemoveByID(db string, collection string, id string) error {
	s, c := connect(db, collection)
	defer s.Close()

	return c.RemoveId(bson.ObjectIdHex(id))
}
