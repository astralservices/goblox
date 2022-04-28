package users

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/astralservices/goblox/goblox/groups"
	"github.com/astralservices/goblox/goblox/network"
)

type UsersHandler struct {
	fetchById       func(id int64) (user IUser, err error)
	fetchByUsername func(username string) (user IUser, err error)
}

type Users struct {
	*network.NetworkRequest
}

func New() *UsersHandler {
	u := &UsersHandler{}

	u.fetchById = func(id int64) (user IUser, err error) {
		ref := &Users{}
		return ref.GetUserById(int(id))
	}

	u.fetchByUsername = func(username string) (user IUser, err error) {
		ref := &Users{}
		return ref.GetUserByUsername(username)
	}

	return u
}

func (ref *Users) GetUserById(userId int) (IUser, error) {
	ref.SetContentType(network.APPJSON)
	ref.SetRequestType(network.GET)
	read, err := ref.SendRequest("https://users.roblox.com/v1/users/"+strconv.Itoa(userId), map[string]interface{}{})
	if err != nil {
		return IUser{}, err
	}

	var r IUser
	err = json.Unmarshal([]byte(read), &r)

	return r, err
}

func (ref *Users) GetUserByUsername(username string) (IUser, error) {
	ref.SetContentType(network.APPJSON)
	ref.SetRequestType(network.POST)
	read, err := ref.SendRequest("https://users.roblox.com/v1/usernames/users", map[string]interface{}{
		"usernames":          []string{username},
		"excludeBannedUsers": false,
	})
	if err != nil {
		return IUser{}, err
	}

	var r IUserByUsername
	err = json.Unmarshal([]byte(read), &r)

	log.Println(r.Data[0].ID)

	u, uErr := ref.GetUserById(int(r.Data[0].ID))

	if uErr != nil {
		return IUser{}, uErr
	}

	return u, err
}

func (ref *Users) GetUserAndPopulate(params UserParams) (IUser, error) {
	ref.SetContentType(network.APPJSON)
	ref.SetRequestType(network.GET)
	read, err := ref.SendRequest("https://users.roblox.com/v1/users/"+strconv.Itoa(params.id), map[string]interface{}{})
	if err != nil {
		return IUser{}, err
	}

	var r IUser
	err = json.Unmarshal([]byte(read), &r)

	if params.populate != nil {
		if params.populate.groups {
			groups, err := ref.GetGroupsForUser(params.id)
			if err != nil {
				return IUser{}, err
			}

			r.groups = &groups
		}
	}

	return r, err
}

func (ref *Users) GetGroupsForUser(userId int) ([]groups.IGroup, error) {
	ref.SetContentType(network.APPJSON)
	ref.SetRequestType(network.GET)
	read, err := ref.SendRequest("https://groups.roblox.com/v1/users/"+strconv.Itoa(userId)+"/groups/roles", map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	var r groups.IUserGroups
	err = json.Unmarshal([]byte(read), &r)

	var groups []groups.IGroup

	for _, v := range r.Data {
		groups = append(groups, v.Group)
	}

	return groups, err
}
