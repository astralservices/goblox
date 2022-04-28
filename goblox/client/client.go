package client

import (
	"net/http"

	"github.com/astralservices/goblox/goblox/network"
	"github.com/astralservices/goblox/goblox/users"
)

type Client struct {
	token string
	http  network.NetworkRequest

	user IAuthenticatedUser

	users users.Users

	login func() bool
}

type Option func(*Client)

func SetToken(token string) Option {
	return func(c *Client) {
		c.token = token
	}
}

func New(opts ...Option) *Client {
	c := &Client{
		http: network.NetworkRequest{},
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

	return c
}

func Login(client *Client) bool {
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
