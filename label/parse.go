package label

import "strconv"

var sku Sku

func ParseOrder(params []string, order *Order) error {
	for i, p := range params {
		switch i {
		case 0:
			order.OrderSn = p
		case 1:
			order.Email = p
		case 2:
			w, err := strconv.Atoi(p)
			if err != nil {
				return err
			}
			order.Weight = w
		case 3:
			order.Consignee = p
		case 4:
			order.Tel = p
		case 5:
			order.State = p
		case 6:
			order.City = p
		case 7:
			order.Street = p
		case 8:
			order.ZipCode = p
		default:
		}
	}
	order.Depot = "001"
	order.Subject = "YKS"
	order.Platform = "WM"
	order.Remark = "DHL_TEST"
	order.Shipping = "DHL SMARTMAIL"
	order.CountryCode = "US"
	order.Money = 50
	sku.Attr = "107"
	order.GoodsList = []Sku{sku}
	return nil
}
