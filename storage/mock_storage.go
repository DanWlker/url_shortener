package storage

import "fmt"

type MockStorage struct {
	data []string
}

func (m *MockStorage) Insert(url string) (id int64, err error) {
	m.data = append(m.data, url)
	return int64(len(m.data) - 1), nil
}

func (m *MockStorage) Retrieve(id int64) (url string, err error) {
	idConverted := int(id)
	if idConverted > len(m.data) {
		return "", IdNotExistError
	}
	return m.data[idConverted], err
}

func (m *MockStorage) Ping() error {
	return nil
}

func NewMockStorage() *MockStorage {
	fmt.Println("Using mock storage client")
	return &MockStorage{}
}
