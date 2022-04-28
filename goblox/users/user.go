package users

import (
	"github.com/astralservices/goblox/goblox/network"
)

type User struct {
	IUser

	network.NetworkRequest
}

func GetGroups(ref *User) {
	ref.SetContentType(network.APPJSON)
	ref.SetRequestType(network.GET)
	ref.SendRequest("https://users.roblox.com/v1/groups", map[string]interface{}{})
}
