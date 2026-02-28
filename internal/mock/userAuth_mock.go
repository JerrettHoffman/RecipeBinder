package mock

import (
	"RecipeBinder/internal"
	"errors"
)

type dbData struct {
	id int
	hashedPassword string
}

type MockUserAuth struct {
	users map[string]dbData
	currentId int
}

func (mock *MockUserAuth) ReadAuthUser(userName string) (internal.UserAuthData, error) {
	data := mock.users[userName]

	if data.id == 0 {
		return internal.UserAuthData{}, errors.New("User not found")
	} else {
		return internal.UserAuthData{
			Id:             data.id,
			UserName:       userName,
			HashedPassword: data.hashedPassword,
		}, nil
	}
}

func (mock *MockUserAuth) CreateAuthUser(userName string, hashedPw string) error {
	if mock.users == nil {
		mock.users = make(map[string]dbData)
	}

	// Incremented first to make the first user 1
	mock.currentId++

	mock.users[userName] = dbData{
		id:             mock.currentId,
		hashedPassword: hashedPw,
	}

	return nil
}

func (mock *MockUserAuth) UpdateAuthUser(currUserId internal.ID, newUser internal.UserAuthData) error {
	return nil
}
