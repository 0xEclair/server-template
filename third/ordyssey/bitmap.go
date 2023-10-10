package ordyssey

import (
	"encoding/json"
	"os"

	"github.com/go-resty/resty/v2"
)

func AllBitmaps() ([]Bitmap, error) {
	u := Base + "/dex/listings/bitmap"
	resp, err := resty.New().R().
		SetAuthToken(os.Getenv("ordyssey")).
		Get(u)
	if err != nil {
		return nil, err
	}

	var res BitmapResult
	err = json.Unmarshal(resp.Body(), &res)
	if err != nil {
		return nil, err
	}

	return res.Results, nil
}
