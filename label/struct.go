// Package label用于从海外物流优选批量获取指定物流渠道追踪码面单
package label

const (
	ApiUrl      = "http://logistics.kokoerp.com/api/?c=api_shipping&a=getLabel"
	LabelPrefix = "http://logistics.kokoerp.com"
)

type Order struct {
	OrderSn      string  `json:"orderSn"`
	Depot        string  `json:"depot"`
	Shipping     string  `json:"shipping"`
	Remark       string  `json:"remark"`
	Subject      string  `json:"subject"`
	SalesAccount string  `json:"salesAccount"`
	Email        string  `json:"email"`
	Amount       float64 `json:"amount"`
	TotalPrice   float64 `json:"totalPrice"`
	Weight       int     `json:"weight"`
	L            int     `json:"l"`
	W            int     `json:"w"`
	H            int     `json:"h"`
	CountryCode  string  `json:"countryCode"`
	State        string  `json:"state"`
	Platform     string  `json:"platform"`
	Consignee    string  `json:"consignee"`
	Tel          string  `json:"tel"`
	Country      string  `json:"country"`
	City         string  `json:"city"`
	Street       string  `json:"street"`
	ZipCode      string  `json:"zipCode"`
	Money        int     `json:"money"`
	Currency     string  `json:"currency"`
	Service      string  `json:"service"`
	OrderCount   int     `json:"orderCount"`
	GoodsList    []Sku   `json:"goodsList"`
}

type Sku struct {
	Sku      string `json:"sku"`
	Quantity int    `json:"quantity"`
	SkuPlace string
	SkuName  string `json:"skuName"`
	Price    float64
	Weight   int
	L        int
	W        int
	H        int
	Attr     string `json:"attr"`
}

type Result struct {
	State int
	Data  Data
	Info  string
}

type Data struct {
	BillData  string
	BillData2 string
	TrackNo   string
}
