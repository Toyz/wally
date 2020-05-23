/*
 * Copyright (c) 2020. BlizzTrack
 */

package mongo_helpers

import (
	"errors"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"log"

)

type MongoSettings struct {
	Host, Username, Password, Database string
}

type Mongo struct {
	session  *mgo.Session
	settings MongoSettings
}

var (
	instance *Mongo
)

func New(settings MongoSettings) *Mongo {
	session, err := mgo.Dial(settings.Host)
	if err != nil {
		// This should never happen but if it does we need to panic... this can cause some wonky effects if we don't...
		log.Panicf("mongodb dial error: %v", err)
	}
	session.SetMode(mgo.Eventual, true)
	session.SetPoolLimit(100)

	if len(settings.Username) > 0 && len(settings.Password) > 0 {
		err := session.Login(&mgo.Credential{
			Username: settings.Username,
			Password: settings.Password,
		})

		if err != nil {
			// Panic when we failed to login because well... go build in logger has no warning...
			// Maybe i should replace the built in logger later... Iris has one built in that we could make public
			log.Panicf("mongodb login failed: %v", err)
		}
	}

	instance = &Mongo{
		session:  session,
		settings: settings,
	}

	return instance
}

func (mg *Mongo) copySession() *mgo.Session {
	return mg.session.Copy()
}

func (mg *Mongo) Collection(collection string) (*mgo.Session, *mgo.Collection) {
	session := mg.copySession()
	c := session.DB(mg.settings.Database).C(collection)

	return session, c
}

func Instance() (*Mongo, error) {
	if instance == nil {
		return nil, errors.New("mongo not configured please run new()")
	}
	return instance, nil
}

func Get(table string) (*mgo.Session, *mgo.Collection) {
	db, err := Instance()
	if err != nil {
		log.Panic(err)
	}

	return db.Collection(table)
}

func One(table string, m bson.M, out interface{}) error {
	session, c := Get(table)
	defer session.Close()

	return c.Find(m).One(out)
}

func OnePipe(table string, m []bson.M, out interface{}) error {
	session, c := Get(table)
	defer session.Close()

	return c.Pipe(m).One(out)
}

func All(table string, m bson.M, out interface{}) error {
	session, c := Get(table)
	defer session.Close()

	return c.Find(m).All(out)
}

func AllPipe(table string, m []bson.M, out interface{}) error {
	session, c := Get(table)
	defer session.Close()

	return c.Pipe(m).All(out)
}

func Insert(table string, input ...interface{}) error {
	session, c := Get(table)
	defer session.Close()

	return c.Insert(input...)
}

func Update(table string, query bson.M, input interface{}) error {
	session, c := Get(table)
	defer session.Close()

	return c.Update(query, input)
}

func Upsert(table string, query bson.M, input interface{}) error {
	session, c := Get(table)
	defer session.Close()

	_, err := c.Upsert(query, input)
	return err
}

func Collection(collection string) (*mgo.Session, *mgo.Collection) {
	db, err := Instance()
	if err != nil {
		log.Panic(err)
	}

	return db.Collection(collection)
}

func Close() error {
	db, err := Instance()
	if err != nil {
		return err
	}

	db.session.Close()
	return nil
}