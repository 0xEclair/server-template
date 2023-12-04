package model

type Asset struct {
	Id            int64
	InscriptionId string
	Address       string
	Type          string
	Category      string
	Collection    string
	Tag           string
	Cons          bool
}

type AssetWithOss struct {
	Asset
	OssUrl string
}

type BRC420Detail struct {
	Id              int64  `gorm:"index"`
	InscriptionId   string `gorm:"index"`
	Address         string `gorm:"index"`
	Tick            string `gorm:"primaryKey;index"`
	Name            string
	Max             string
	Rest            string `gorm:"index"`
	Price           string
	DeployTimestamp int32
	RawTick         string `gorm:"index"`
}

type BRC420Entry struct {
	Id            int64  `gorm:"primaryKey;index"`
	InscriptionId string `gorm:"primaryKey;index"`
	Address       string `gorm:"index"`
	ContentType   string `gorm:"index"`
	Ref           string `gorm:"index"`
	FeeAt         string `gorm:"index"`
	MintTimestamp int32
}

type BRC420EntryWithName struct {
	Id            int64
	InscriptionId string
	Address       string
	ContentType   string
	Ref           string
	FeeAt         string
	MintTimestamp int32
	Name          string
}

type BRC420EntryWithTick struct {
	BRC420Entry
	Tick string
}
