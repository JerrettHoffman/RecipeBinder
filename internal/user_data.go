package internal

type UserAuthData struct {
	Id             ID
	UserName       string
	HashedPassword string
}

type UserAuthDataStrategy interface {
	ReadAuthUser(userName string) (UserAuthData, error)
	CreateAuthUser(userName string, hashedPw string) error

	// more though needed here for the forgot password scenario
	UpdateAuthUser(currUserId ID, newUser UserAuthData) error
}
