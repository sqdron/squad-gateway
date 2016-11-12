package remote

import (
	"github.com/pkg/errors"
	"github.com/sqdron/squad"
	"fmt"
)

type remoteAuth struct {
	api squad.ISquadAPI
}

type IRemoteAuth interface {
	Signup(email string) error
	Login(user string, password string) (string, error)
	Validate(token string) error
}

func RemoteAuth(api squad.ISquadAPI) *remoteAuth {
	return &remoteAuth{api:api}
}

type SignupRequest struct {
	Email string
}

func (c *remoteAuth) Signup(email string) error {
	fmt.Println("email", email)
	_, err := c.api.Request("passport.signup", &SignupRequest{Email:email})
	return err
}

type LoginRequest struct {
	User     string
	Password string
}

func (c *remoteAuth) Login(user string, password string) (string, error) {
	res, err := c.api.Request("passport.login", &LoginRequest{User:user, Password:password})

	return string(res.([]byte)), err
}

type TokenRequest struct {
	Token string
}

func (c *remoteAuth) Validate(token string) error {
	if (token == "") {
		return errors.New("Token is empty")
	}

	_, err := c.api.Request("passport.validate", &TokenRequest{Token:token})
	fmt.Println("Validate res", err)
	return err
}