package goblox

import (
	"encoding/json"
	"strconv"
)

type User struct {
	IUser
	client *Client
}

func GetGroups(ref *Client) {
	ref.http.SetContentType(APPJSON)
	ref.http.SetRequestType(GET)
	ref.http.SendRequest("https://users.roblox.com/v1/groups", map[string]interface{}{})
}

func (ref *User) GetUsernameHistory() (IPagedResponse[IUsernameHistory], error) {
	ref.client.http.SetContentType(APPJSON)
	ref.client.http.SetRequestType(GET)
	data, err := ref.client.http.SendRequest("https://users.roblox.com/v1/users/"+strconv.Itoa(int(ref.ID))+"/username-history", map[string]interface{}{})

	if err != nil {
		return IPagedResponse[IUsernameHistory]{}, err
	}

	var r IPagedResponse[IUsernameHistory]
	err = json.Unmarshal([]byte(data), &r)

	return r, err
}

func (ref *User) GetGroups() (IUserGroups, error) {
	ref.client.http.SetContentType(APPJSON)
	ref.client.http.SetRequestType(GET)
	data, err := ref.client.http.SendRequest("https://groups.roblox.com/v1/users/"+strconv.Itoa(int(ref.ID))+"/groups/roles", map[string]interface{}{})

	if err != nil {
		return IUserGroups{}, err
	}
	var r IUserGroups
	err = json.Unmarshal([]byte(data), &r)

	return r, err
}

func (ref *User) New(data *IUser, client Client) *User {
	user := &User{
		IUser: *data,

		client: &client,
	}

	return user
}
