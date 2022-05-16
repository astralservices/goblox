package goblox

import (
	"encoding/json"
	"net/http"
)

type Client struct {
	Token string

	User IAuthenticatedUser

	Users  *UsersHandler
	Groups *GroupsHandler

	Login func() bool

	http NetworkRequest
}

type Option func(Client)

// Sets the token for the client
//
// Include the full token, including the warning prefix.
func SetToken(token string) Option {
	return func(c Client) {
		c.Token = token
	}
}

// Creates a new client with options
func New(opts ...Option) *Client {
	c := Client{
		http: NetworkRequest{},
	}

	c.http.New()

	for _, opt := range opts {
		opt(c)
	}

	if c.Token != "" {
		c.http.AddCookie(&http.Cookie{
			Name:  ".ROBLOSECURITY",
			Value: c.Token,
		})
	}

	c.Login = func() bool {
		return Login(c)
	}

	c.Users = NewUsersHandler(c)

	c.Groups = NewGroupsHandler(c)

	return &c
}

func Login(client Client) bool {
	user, err := GetCurrentUser(client)

	if err != nil {
		return false
	}

	authed := user.ID != 0

	if authed {
		client.User = *user
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
