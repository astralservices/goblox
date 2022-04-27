package users

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/astralservices/goblox/goblox/groups"
	"github.com/astralservices/goblox/goblox/network"
)

func (ref *User) GetUserById(userId int) (IUser, error) {
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

func (ref *User) GetUserByUsername(username string) (IUser, error) {
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

func (ref *User) GetGroupsForUser(userId int) ([]groups.IGroup, error) {
	ref.SetContentType(network.APPJSON)
	ref.SetRequestType(network.GET)
	read, err := ref.SendRequest("https://users.roblox.com/v1/users/"+strconv.Itoa(userId)+"/groups", map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	var r groups.IUserGroups
	err = json.Unmarshal([]byte(read), &r)
}
