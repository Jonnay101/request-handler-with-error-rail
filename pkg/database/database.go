package database

import (
	"io/ioutil"
	"os"

	"github.com/music-tribe/uuid"
)

// request must be staisfied to use the db methods
type request interface {
	GetFilename() string
	GetFileData() []byte
	GetID() uuid.UUID
}

// Database shows the required methods to satisfy this project
type Database interface {
	Post(request) error
}

// ImageStore is this packages db
type imageStore struct{}

// CreateNewImageStore will create a new database session
func CreateNewImageStore() (Database, error) {
	return &imageStore{}, nil
}

// Post will add an image to the db
func (s *imageStore) Post(r request) error {
	if err := os.MkdirAll("imageStore", os.ModePerm); err != nil {
		return err
	}
	ioutil.WriteFile(r.GetFilename(), r.GetFileData(), os.ModePerm)
	return nil
}
