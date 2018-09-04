package database

import (
	scribble "github.com/nanobox-io/golang-scribble"
)

// New creates a new database
func New(dir string) (*scribble.Driver, error) {
	db, err := scribble.New(dir, nil)
	if err != nil {
		return nil, err
	}
	return db, nil
}
