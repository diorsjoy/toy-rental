package models

import (
	"errors"
	"time"
)

var (
	ErrNoRecord = errors.New("models: no matching record found")

	// ErrInvalidCredentials error. We'll use this later if a user
	// tries to login with an incorrect email address or password.
	ErrInvalidCredentials = errors.New("models: invalid credentials")

	// ErrDuplicateEmail error. We'll use this later if a user
	// tries to signup with an email address that's already in use.
	ErrDuplicateEmail = errors.New("models: duplicate email")

	// ErrEmailDoesNotExist error
	ErrEmailDoesNotExist = errors.New("models: email does not exist")
)

type Feedback struct {
	ID      int
	Name    string
	Content string
	Stars   int
}

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Role           string
	Created        time.Time
	Active         bool
}
