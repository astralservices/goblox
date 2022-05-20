package goblox

import (
	"encoding/json"
	"strconv"
)

type MarketplaceHandler struct {
	client *Client
}

// Creates a new user handler with the given client.
//
// A user handler is used to fetch users by ID and username.
func NewMarketplaceHandler(client Client) *MarketplaceHandler {
	return &MarketplaceHandler{
		client: &client,
	}
}

func (ref *MarketplaceHandler) GetProductById(productId int64) (*IProduct, error) {
	ref.client.http.SetContentType(APPJSON)
	ref.client.http.SetRequestType(GET)
	read, err := ref.client.http.SendRequest("https://api.roblox.com/marketplace/productinfo?assetId="+strconv.Itoa(int(productId)), map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	var r IProduct
	err = json.Unmarshal([]byte(read), &r)

	return &r, err
}

func (ref *MarketplaceHandler) GetResellersForProduct(productId int64) ([]IReseller, error) {
	ref.client.http.SetContentType(APPJSON)
	ref.client.http.SetRequestType(GET)
	read, err := ref.client.http.SendRequest("https://economy.roblox.com/v1/assets/" + strconv.Itoa(int(productId)) + "/resellers?cursor=&limit=100", map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	var r IPagedResponse[IReseller]
	err = json.Unmarshal([]byte(read), &r)

	return r.Data, err
}

func (ref *MarketplaceHandler) BuyResoldProduct(productId int64, product IReseller) (bool, error) {
	ref.client.http.SetContentType(APPJSON)
	ref.client.http.SetRequestType(POST)
	read, err := ref.client.http.SendRequest("https://economy.roblox.com/v1/purchases/products/"+strconv.Itoa(int(productId)), map[string]interface{}{
		"expectedCurrency": 1,
		"userAssetID": product.UserAssetID,
		"isLimited": true,
		"expectedPrice": product.Price,
		"productId": productId,
		"expectedSellerId": product.Seller.ID,
	})

	if err != nil {
		return false, err
	}

	var r IResellerBuyResponse

	err = json.Unmarshal([]byte(read), &r)

	return r.Purchased, err
}