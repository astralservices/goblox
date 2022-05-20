package goblox

type IProduct struct {
	TargetID               int64           `json:"TargetId"`
	ProductType            string          `json:"ProductType"`
	AssetID                int64           `json:"AssetId"`
	ProductID              int64           `json:"ProductId"`
	Name                   string          `json:"Name"`
	Description            string          `json:"Description"`
	AssetTypeID            int64           `json:"AssetTypeId"`
	Creator                IProductCreator `json:"Creator"`
	IconImageAssetID       int64           `json:"IconImageAssetId"`
	Created                string          `json:"Created"`
	Updated                string          `json:"Updated"`
	PriceInRobux           int64           `json:"PriceInRobux"`
	PriceInTickets         interface{}     `json:"PriceInTickets"`
	Sales                  int64           `json:"Sales"`
	IsNew                  bool            `json:"IsNew"`
	IsForSale              bool            `json:"IsForSale"`
	IsPublicDomain         bool            `json:"IsPublicDomain"`
	IsLimited              bool            `json:"IsLimited"`
	IsLimitedUnique        bool            `json:"IsLimitedUnique"`
	Remaining              interface{}     `json:"Remaining"`
	MinimumMembershipLevel int64           `json:"MinimumMembershipLevel"`
}

type IProductCreator struct {
	ID   int64  `json:"Id"`
	Name string `json:"Name"`
}

type IReseller struct {
	UserAssetID  int64       `json:"userAssetId"`
	Seller       ISeller     `json:"seller"`
	Price        int64       `json:"price"`
	SerialNumber interface{} `json:"serialNumber"`
}

type ISeller struct {
	ID   int64        `json:"id"`
	Type ResellerType `json:"type"`
	Name string       `json:"name"`
}

type ResellerType string

const (
	ResellerUserType ResellerType = "User"
)

type IResellerBuyResponse struct {
	Purchased        bool   `json:"purchased"`
	Reason           string `json:"reason"`
	ProductID        int64  `json:"productId"`
	StatusCode       int64  `json:"statusCode"`
	Title            string `json:"title"`
	ErrorMsg         string `json:"errorMsg"`
	ShowDivID        string `json:"showDivId"`
	ShortfallPrice   int64  `json:"shortfallPrice"`
	BalanceAfterSale int64  `json:"balanceAfterSale"`
	ExpectedPrice    int64  `json:"expectedPrice"`
	Currency         int64  `json:"currency"`
	Price            int64  `json:"price"`
	AssetID          int64  `json:"assetId"`
}
