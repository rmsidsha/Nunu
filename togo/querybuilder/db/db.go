package db

import (
	"errors"
	"sync"
)

type DB interface {
	Table(name string)     // set table name
	Query(q string) string // perform raw query on the table
	Find(id int) string    // find a record by it's primary id
	Get() string           // fetch all the records from the table
}

const (
	_ = iota
	MYSQL
)

var db DB          // will be used as a asingleton db object
var once sync.Once // make thread safe singleton

// Return a db singleton object
func New(driver int) (DB, error) {
	var err error
	once.Do(func() {
		db, err = databaseFactory(driver)
	})
	return db, err
}

// Database factory
// Return a DB for general purpose
func databaseFactory(driver int) (DB, error) {
	if driver == MYSQL {
		return new(Mysql), nil
	}
	return nil, errors.New("Unsupported storage driver")
}
