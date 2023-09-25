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

type BitmapListInfo struct {
	Id            int64  `gorm:"primaryKeym,index"`
	InscriptionId string `gorm:"index"`
	ContentType   string `gorm:"index"`
	Address       string `gorm:"index"`
	Listed        bool   `gorm:"index"`
	ListedPrice   uint64 `gorm:"index"`
	Domain        string `gorm:"index"`
	BitmapId      int64  `gorm:"index"`
	Name          string
	Content       string `gorm:"index"`
	Source        string `gorm:"index"`
}
