package mockrepo

import "errors"

var mapDatabase = map[int]string{}

type MockRepo struct {
}

func NewMockRepo() *MockRepo {
	return &MockRepo{}
}

func (m MockRepo) Get(id int) (string, error) {

	value, ok := mapDatabase[id]
	if !ok {
		return "", nil
	}

	return value, nil
}

func (m MockRepo) Set(clientID string) error {

	_, ok := mapDatabase[1]
	if ok {
		return errors.New("already exists")
	}

	mapDatabase[1] = clientID

	return nil
}
