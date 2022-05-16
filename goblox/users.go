package goblox

import (
	"encoding/json"
	"strconv"
)

type UsersHandler struct {
	client *Client
}

// Creates a new user handler with the given client.
//
// A user handler is used to fetch users by ID and username.
func NewUsersHandler(client Client) *UsersHandler {
	return &UsersHandler{
		client: &client,
	}
}

// Gets a user by ID.
func (ref *UsersHandler) GetUserById(userId int64) (*User, error) {
	ref.client.http.SetContentType(APPJSON)
	ref.client.http.SetRequestType(GET)
	read, err := ref.client.http.SendRequest("https://users.roblox.com/v1/users/"+strconv.Itoa(int(userId)), map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	var r IUser
	err = json.Unmarshal([]byte(read), &r)

	user := User{
		IUser:  r,
		client: ref.client,
	}

	u := user.New(&user.IUser, *user.client)

	return u, err
}

// Gets a user by username.
func (ref *UsersHandler) GetUserByUsername(username string) (*User, error) {
	ref.client.http.SetContentType(APPJSON)
	ref.client.http.SetRequestType(POST)
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

// Gets a user by ID and populates specified fields.
func (ref *UsersHandler) GetUserAndPopulate(params UserParams) (User, error) {
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

// Gets all the groups for a user by ID.
func (ref *UsersHandler) GetGroupsForUser(userId int) ([]IGroup, error) {
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
