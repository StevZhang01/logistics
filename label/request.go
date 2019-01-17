package label

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func Request(orderData string) (*Result, error) {
	values := url.Values{"data": {orderData}}
	resp, err := http.PostForm(ApiUrl, values)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Request failed: %s", resp.Status)
	}
	var result Result
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil

}
