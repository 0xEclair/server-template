package bestinslot

type SatsResponse struct {
	Data        map[string]string `json:"data"`
	BlockHeight int32             `json:"block_height"`
}

type BRC20Response struct {
	Data        []BRC20 `json:"data"`
	BlockHeight int32   `json:"block_height"`
}

type BRC20 struct {
	Ticker              string `json:"ticker"`
	OverallBalance      string `json:"overall_balance"`
	AvaliableBalance    string `json:"available_balance"`
	TransferableBalance string `json:"transferrable_balance"`
}
