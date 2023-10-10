package serializer

import (
	"server-template/model"
	"server-template/third/ordyssey"
)

type OrdysseyBitmapResponse struct {
	Id            int64  `json:"id"`
	InscriptionId string `json:"inscription_id"`
	ContentType   string `json:"content_type"`
	Address       string `json:"address"`
	Listed        bool   `json:"listed"`
	Price         uint64 `json:"list_price"`
	BitmapId      int64  `json:"bitmap_id"`
	Content       string `json:"content"`
	Source        string `json:"source"`
	SourceUrl     string `json:"source_url"`
}

func BuildOrdysseyBitmapResponse(b model.Bitmap, o ordyssey.Bitmap) OrdysseyBitmapResponse {
	return OrdysseyBitmapResponse{
		Id:            b.Id,
		InscriptionId: b.InscriptionId,
		BitmapId:      int64(b.BitmapId),
		Price:         o.Price,
		Source:        o.Source,
		SourceUrl:     SOURCE_URL[o.Source],
	}
}

type OrdysseyBitmapListResponse struct {
	Cnt   uint64                   `json:"count"`
	Items []OrdysseyBitmapResponse `json:"items"`
}

func BuildOrdysseyBitmapListResponse(bs []model.Bitmap, os map[string]ordyssey.Bitmap) OrdysseyBitmapListResponse {
	var list OrdysseyBitmapListResponse

	for _, bitmap := range bs {
		r := BuildOrdysseyBitmapResponse(bitmap, os[bitmap.InscriptionId])
		list.Items = append(list.Items, r)
	}

	list.Cnt = uint64(len(bs))

	return list
}
