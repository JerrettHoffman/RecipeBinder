package auth

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserData struct {
	hashedPassword []byte
	Id             int
}

var currentId = 1
var userDatabase map[string]UserData
var sessionManager *scs.SessionManager

func Setup() {
	sessionManager = scs.New()
	sessionManager.Lifetime = 12 * time.Hour

	userDatabase = make(map[string]UserData, 0)

	currentId = 1
}

func CreateUser(name string, password string) error {
	if hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12); err != nil {
		return err
	} else {
		userDatabase[name] = UserData{hashedPassword, currentId}
		currentId++
		return nil
	}
}

func Authenticate(name string, password string, ctx context.Context) error {
	// Hit the db
	userData := userDatabase[name]

	if err := bcrypt.CompareHashAndPassword(userData.hashedPassword, []byte(password)); err != nil {
		return err
	} else {
		sessionManager.Put(ctx, "user", name)
		sessionManager.Put(ctx, "id", userData.Id)
		return nil
	}
}

func GetUserId(ctx context.Context) (int, error) {
	if id := sessionManager.GetInt(ctx, "id"); id == 0 {
		return 0, errors.New("Failed to parse Id")
	} else {
		return id, nil
	}
}

func SessionMiddleware(handler http.Handler) http.Handler {
	return sessionManager.LoadAndSave(handler)
}
