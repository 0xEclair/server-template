package model

type Asset struct {
	Id            int64
	InscriptionId string
	Address       string
	Type          string
	Category      string
}

type AssetWithOss struct {
	Asset
	OssUrl string
}
