package database

import (
	scribble "github.com/nanobox-io/golang-scribble"
)

type Database interface {
	Delete(string, string) error
	Read(string, string, interface{}) error
	ReadAll(string) ([]string, error)
	Write(string, string, interface{}) error
}
type database struct {
	*scribble.Driver
}

// New creates a new database
func New(dir string) (Database, error) {
	db, err := scribble.New(dir, nil)
	if err != nil {
		return nil, err
	}
	return db, nil
}
