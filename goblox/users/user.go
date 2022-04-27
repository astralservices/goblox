package users

import "github.com/astralservices/goblox/goblox/network"

func GetGroups(ref *User) {
	ref.SetContentType(network.APPJSON)
	ref.SetRequestType(network.GET)
	ref.SendRequest("https://users.roblox.com/v1/groups", map[string]interface{}{})
}
