package bestinslot

import (
	"encoding/json"
	"net/url"
	"os"

	"github.com/go-resty/resty/v2"
)

func AddressBySats(sats string) (string, error) {
	u := Base + "/sats/forward_lookup?sats_name=" + url.QueryEscape(sats)
	resp, err := resty.New().R().
		SetHeader("x-api-key", os.Getenv("bestinslot")).
		Get(u)
	if err != nil {
		return "", err
	}

	var res SatsResponse
	err = json.Unmarshal(resp.Body(), &res)
	if err != nil {
		return "", err
	}

	return res.Data["wallet"], nil
}
