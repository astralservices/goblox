package marketplace

import (
	"log"
	"testing"

	"github.com/astralservices/goblox/goblox/network"
)

func Test_GetProductInfo(test *testing.T) {
	market := Marketplace{}
	market.New()
	market.SetRequestType(network.GET)
	market.SetContentType(network.APPJSON)
	list, err := market.GetProductInfo(1584277735)
	if err != nil {
		test.Errorf(err.Error())
	}
	log.Println(list)
}
