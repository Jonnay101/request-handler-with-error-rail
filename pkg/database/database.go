package database

import (
	"fmt"
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

// Database - these methods are required to satisfy this interface
type Database interface {
	Post(request) error
}

// ImageStore is this packages db
type imageStore struct {
	StoragePath string
}

// CreateNewImageStore will create a new database session
func CreateNewImageStore(storagePath string) Database {
	return &imageStore{StoragePath: storagePath}
}

// createStorageKeyPath uses the existing db storage path and the images uuid to create a unique storage path
func (s *imageStore) createStorageKeyPath(imageUUID uuid.UUID) string {
	imgUUIDString := imageUUID.String()
	return fmt.Sprintf("%s/%s/%s", s.StoragePath, imgUUIDString[:2], imgUUIDString)
}

// Post will add an image to the db
func (s *imageStore) Post(r request) error {
	if err := os.MkdirAll("imageStore", os.ModePerm); err != nil {
		return err
	}
	ioutil.WriteFile(r.GetFilename(), r.GetFileData(), os.ModePerm)
	return nil
}
