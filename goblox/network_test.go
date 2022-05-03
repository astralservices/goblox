package goblox

import (
	"log"
	"testing"
)

// Test_main :
func Test_NetworkRequest(test *testing.T) {
	req := NetworkRequest{}
	req.New()

	req.SetContentType(APPJSON)
	req.SetRequestType(POST)
	log.Println(req.SendRequest("https://users.roblox.com/v1/usernames/users", map[string]interface{}{
		"usernames":          []string{"AstralServices"},
		"excludeBannedUsers": true,
	}))
}
