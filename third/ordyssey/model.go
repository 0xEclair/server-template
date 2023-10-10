package ordyssey

type BitmapResult struct {
	Results []Bitmap `json:"results"`
}

type Bitmap struct {
	Price         uint64 `json:"price"`
	Id            int64  `json:"id"`
	Source        string `json:"source"`
	InscriptionId string `json:"_id"`
}
