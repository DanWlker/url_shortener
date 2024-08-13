package storage

import "errors"

var IdNotExistError = errors.New("Id requested is not associated with any url")

type StorageClient interface {
	Insert(url string) (id int64, err error)
	Retrieve(id int64) (url string, err error)
	Ping() error
}
