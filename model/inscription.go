package model

type Inscription struct {
	Id            int64
	InscriptionId string
	Address       string
	Content       string
	ContentType   string
	OssUrl        string
}

type InscriptionWithNameCount struct {
	Id            int64
	InscriptionId string
	Address       string
	Content       string
	ContentType   string
	OssUrl        string
	Name          string
	Count         int64
}
