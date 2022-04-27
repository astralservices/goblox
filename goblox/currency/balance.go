// api.roblox.com/currency

package currency

import (
	"log"

	"github.com/astralservices/goblox/goblox/authenticate"
)

func (ref *Currency) GetBalance(auth authenticate.Authenticate) (map[string]interface{}, error) {
	str, err := ref.SendRequest("http://api.roblox.com/currency/balance", map[string]interface{}{})
	if err == nil {
		return map[string]interface{}{}, err
	}
	log.Println(str)
	return map[string]interface{}{}, nil
}
