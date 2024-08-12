package storage

type StorageClient interface {
	Insert(url string) (id int64, err error)
	Retrieve(id int64) (url string, err error)
}
