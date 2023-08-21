package model

type Bitmap struct {
	Id            int64
	InscriptionId string
	Address       string
	BitmapId      int32
	GenesisHeight int32
	Timestamp     int32
}

type BitmapRank struct {
	Address string `json:"address"`
	Number  int32  `json:"number"`
}
