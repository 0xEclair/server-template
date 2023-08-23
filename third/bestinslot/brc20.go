package bestinslot

import (
	"encoding/json"
	"os"

	"github.com/go-resty/resty/v2"
)

func BRC20ByAddress(address string) ([]BRC20, error) {
	resp, err := resty.New().R().
		SetHeader("x-api-key", os.Getenv("bestinslot")).
		Get(Base + "/brc20/wallet_balances?address=" + address)
	if err != nil {
		return nil, err
	}

	var res BRC20Response
	err = json.Unmarshal(resp.Body(), &res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}
