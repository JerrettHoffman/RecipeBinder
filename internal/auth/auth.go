package auth

import (
	"RecipeBinder/internal"
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserData struct {
	Id   int
	User string
}

const (
	UninitialzedId = -1
)

var sessionManager *scs.SessionManager

func Setup() {
	sessionManager = scs.New()
	sessionManager.Lifetime = 12 * time.Hour
}

func CreateUser(name string, password string, userDatabase internal.UserAuthDataStrategy) error {
	if hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12); err != nil {
		return err
	} else {
		return userDatabase.CreateAuthUser(name, string(hashedPassword))
	}
}

func Authenticate(name string, password string, ctx context.Context, userDatabase internal.UserAuthDataStrategy) error {
	if userData, err := userDatabase.ReadAuthUser(name); err != nil {
		return err
	} else {
		if err := bcrypt.CompareHashAndPassword([]byte(userData.HashedPassword), []byte(password)); err != nil {
			return err
		} else {
			sessionManager.Put(ctx, "user", name)
			sessionManager.Put(ctx, "userId", userData.Id)
			return nil
		}
	}
}

func GetUser(ctx context.Context) (UserData, error) {
	if id := sessionManager.GetInt(ctx, "userId"); id == 0 {
		return UserData{UninitialzedId, ""}, errors.New("Failed to parse userId")
	} else {
		if user := sessionManager.GetString(ctx, "user"); user == "" {
			return UserData{UninitialzedId, ""}, errors.New("Failed to parse user")
		} else {
			return UserData{id, user}, nil
		}
	}
}

func SessionMiddleware(handler http.Handler) http.Handler {
	return sessionManager.LoadAndSave(handler)
}
