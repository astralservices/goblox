package goblox

import (
	"encoding/json"
	"log"
	"strconv"
)

type UsersHandler struct {
	fetchById       func(id int64) (user *User, err error)
	fetchByUsername func(username string) (user *User, err error)
}

type Users struct {
	client Client
}

func (ref *UsersHandler) New(client *Client) *UsersHandler {
	u := &UsersHandler{}

	u.fetchById = func(id int64) (user *User, err error) {
		ref := &Users{
			client: *client,
		}
		return ref.GetUserById(int64(id))
	}

	u.fetchByUsername = func(username string) (user *User, err error) {
		ref := &Users{
			client: *client,
		}
		return ref.GetUserByUsername(username)
	}

	return u
}

func (ref *Users) GetUserById(userId int64) (*User, error) {
	ref.client.http.SetContentType(APPJSON)
	ref.client.http.SetRequestType(GET)
	log.Println("sending request")
	read, err := ref.client.http.SendRequest("https://users.roblox.com/v1/users/"+strconv.Itoa(int(userId)), map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	var r IUser
	err = json.Unmarshal([]byte(read), &r)

	user := User{
		IUser:  r,
		client: &ref.client,
	}

	u := user.New(&user.IUser, *user.client)

	return u, err
}

func (ref *Users) GetUserByUsername(username string) (*User, error) {
	ref.client.http.SetContentType(APPJSON)
	ref.client.http.SetRequestType(POST)
	log.Println("sending request")
	read, err := ref.client.http.SendRequest("https://users.roblox.com/v1/usernames/users", map[string]interface{}{
		"usernames":          []string{username},
		"excludeBannedUsers": false,
	})
	if err != nil {
		return nil, err
	}

	var r IUserByUsername
	err = json.Unmarshal([]byte(read), &r)

	u, uErr := ref.GetUserById(int64(r.Data[0].ID))

	if uErr != nil {
		return nil, uErr
	}

	return u, err
}

func (ref *Users) GetUserAndPopulate(params UserParams) (User, error) {
	ref.client.http.SetContentType(APPJSON)
	ref.client.http.SetRequestType(GET)
	read, err := ref.client.http.SendRequest("https://users.roblox.com/v1/users/"+strconv.Itoa(params.id), map[string]interface{}{})
	if err != nil {
		return User{}, err
	}

	var r IUser
	err = json.Unmarshal([]byte(read), &r)

	if params.populate != nil {
		if params.populate.groups {
			groups, err := ref.GetGroupsForUser(params.id)
			if err != nil {
				return User{}, err
			}

			r.groups = &groups
		}
	}

	user := User{
		IUser: r,
	}

	return user, err
}

func (ref *Users) GetGroupsForUser(userId int) ([]IGroup, error) {
	ref.client.http.SetContentType(APPJSON)
	ref.client.http.SetRequestType(GET)
	read, err := ref.client.http.SendRequest("https://groups.roblox.com/v1/users/"+strconv.Itoa(userId)+"/groups/roles", map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	var r IUserGroups
	err = json.Unmarshal([]byte(read), &r)

	var groups []IGroup

	for _, v := range r.Data {
		groups = append(groups, v.Group)
	}

	return groups, err
}
