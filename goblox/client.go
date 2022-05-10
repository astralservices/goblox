package goblox

import (
	"encoding/json"
	"net/http"
)

type Client struct {
	token string
	http  NetworkRequest

	user IAuthenticatedUser

	users  UsersHandler
	groups GroupsHandler

	login func() bool
}

type Option func(Client)

// Sets the token for the client
//
// Include the full token, including the warning prefix.
func SetToken(token string) Option {
	return func(c Client) {
		c.token = token
	}
}

// Creates a new client with options
func New(opts ...Option) Client {
	c := Client{
		http: NetworkRequest{},
	}

	c.http.New()

	for _, opt := range opts {
		opt(c)
	}

	if c.token != "" {
		c.http.AddCookie(&http.Cookie{
			Name:  ".ROBLOSECURITY",
			Value: c.token,
		})
	}

	c.login = func() bool {
		return Login(c)
	}

	c.users = UsersHandler{}

	c.users = *c.users.New(c)

	c.groups = GroupsHandler{}

	c.groups = *c.groups.New(c)

	return c
}

func Login(client Client) bool {
	user, err := GetCurrentUser(client)

	if err != nil {
		return false
	}

	authed := user.ID != 0

	if authed {
		client.user = *user
	}

	return authed
}

func GetCurrentUser(client Client) (user *IAuthenticatedUser, err error) {
	client.http.SetRequestType(GET)
	read, err := client.http.SendRequest("https://users.roblox.com/v1/users/authenticated", map[string]interface{}{})

	if err != nil {
		return &IAuthenticatedUser{}, err
	}

	var r IAuthenticatedUser
	err = json.Unmarshal([]byte(read), &r)

	return &r, err
}
