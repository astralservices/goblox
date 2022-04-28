package client

import (
	"encoding/json"

	"github.com/astralservices/goblox/goblox/network"
)

type IAuthenticatedUser struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}

func GetCurrentUser(client *Client) (user *IAuthenticatedUser, err error) {
	client.http.SetRequestType(network.GET)
	read, err := client.http.SendRequest("https://users.roblox.com/v1/users/authenticated", map[string]interface{}{})

	if err != nil {
		return &IAuthenticatedUser{}, err
	}

	var r IAuthenticatedUser
	err = json.Unmarshal([]byte(read), &r)

	return &r, err
}
