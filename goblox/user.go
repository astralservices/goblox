package goblox

import (
	"encoding/json"
	"strconv"
)

type User struct {
	IUser

	http *NetworkRequest
}

func GetGroups(ref *Client) {
	ref.http.SetContentType(APPJSON)
	ref.http.SetRequestType(GET)
	ref.http.SendRequest("https://users.roblox.com/v1/groups", map[string]interface{}{})
}

func (ref *User) GetUsernameHistory() (IPagedResponse[IUsernameHistory], error) {
	ref.http.SetContentType(APPJSON)
	ref.http.SetRequestType(GET)
	data, err := ref.http.SendRequest("https://users.roblox.com/v1/users"+strconv.Itoa(int(ref.ID))+"/username-history", map[string]interface{}{})

	if err != nil {
		return IPagedResponse[IUsernameHistory]{}, err
	}

	var r IPagedResponse[IUsernameHistory]
	err = json.Unmarshal([]byte(data), &r)

	return r, err
}

func (ref *User) New(data *IUser, http *NetworkRequest) *User {
	user := &User{
		IUser: *data,

		http: http,
	}

	return user
}
