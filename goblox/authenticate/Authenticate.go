// auth.roblox.com
// Authorises Username/Password combination
// Use for requests that require you to be logged in
// ---------------------------------------------------------
// Example usage for calling a request that requires a login:
// This example shows the Currency/Balance API
/*
	auth := authenticate.Authenticate{}
	err := auth.Login("username", "password")
	if err != nil {
		return err
	}
	cur := currency.Currency{}
	cur.New()
	cur.SetHeaders(auth.GetResponse().Header)
	for _, cookie := range auth.GetCookieArray() {
		cur.AddCookie(cookie)
	}
	cur.SetContentType(network.APPJSON)
	cur.SetRequestType(network.GET)
	data, err := cur.SendRequest("http://api.roblox.com/currency/balance", map[string]interface{}{})
	if err != nil {
		test.Errorf(err.Error())
	}
	log.Println(data)

*/

package authenticate

import (
	"log"

	"github.com/astralservices/goblox/goblox/network"
)

type Authenticate struct {
	network.NetworkRequest
}

func (ref *Authenticate) Login(username, password string) error {
	ref.New()
	ref.Logout()
	ref.AddHeader("X-CSRF-TOKEN", []string{ref.GetResponse().Header["X-Csrf-Token"][0]})
	ref.SetContentType(network.APPJSON)
	ref.SetRequestType(network.POST)
	ref.AddHeader("Accept", []string{"application/json"})
	ref.AddHeader("Content-Type", []string{"application/json"})
	str, err := ref.SendRequest("https://auth.roblox.com/v2/login", map[string]interface{}{
		"ctype":    "Username",
		"cvalue":   username,
		"password": password,
	})
	if err != nil {
		return err
	}
	log.Println("Auth: ")
	log.Println(str)
	return nil
}

func (ref *Authenticate) Logout() {
	ref.New()
	ref.SetContentType(network.APPJSON)
	ref.SetRequestType(network.POST)
	ref.SendRequest("https://api.roblox.com/sign-out/v1", map[string]interface{}{})
}
