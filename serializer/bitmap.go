package serializer

import (
	"server-template/model"
)

var SOURCE_URL map[string]string = map[string]string{
	"magic_eden": "https://magiceden.io/ordinals/item-details/",
}

type BitmapResponse struct {
	Id            int64  `json:"id"`
	InscriptionId string `json:"inscription_id"`
	Address       string `json:"address"`
	BitmapId      int32  `json:"bitmap_id"`
	GenesisHeight int32  `json:"genesis_height"`
}

func BuildBitmapResponse(inscription model.Bitmap) *BitmapResponse {
	return &BitmapResponse{
		Id:            inscription.Id,
		InscriptionId: inscription.InscriptionId,
		Address:       inscription.Address,
		BitmapId:      inscription.BitmapId,
		GenesisHeight: inscription.GenesisHeight,
	}
}

type BitmapListResponse struct {
	Count int64             `json:"count"`
	Items []*BitmapResponse `json:"items"`
}

func BuildBitmapListResponse(items []model.Bitmap) []*BitmapResponse {
	var inscriptions []*BitmapResponse
	for _, item := range items {
		inscription := BuildBitmapResponse(item)
		inscriptions = append(inscriptions, inscription)
	}

	return inscriptions
}

func BuildBitmapListWithCntResponse(cnt int64, items []model.Bitmap) *BitmapListResponse {
	var inscriptions []*BitmapResponse
	for _, item := range items {
		inscription := BuildBitmapResponse(item)
		inscriptions = append(inscriptions, inscription)
	}

	return &BitmapListResponse{
		Count: cnt,
		Items: inscriptions,
	}
}

type BitmapListInfoResponse struct {
	Id            int64  `json:"id"`
	InscriptionId string `json:"inscription_id"`
	ContentType   string `json:"content_type"`
	Address       string `json:"address"`
	Listed        bool   `json:"listed"`
	ListedPrice   uint64 `json:"list_price"`
	BitmapId      int64  `json:"bitmap_id"`
	Content       string `json:"content"`
	Source        string `json:"source"`
	SourceUrl     string `json:"source_url"`
}

func BuildBitmapListInfoResponse(bitmap model.BitmapListInfo) BitmapListInfoResponse {

	return BitmapListInfoResponse{
		Id:            bitmap.Id,
		InscriptionId: bitmap.InscriptionId,
		ContentType:   bitmap.ContentType,
		Address:       bitmap.Address,
		Listed:        bitmap.Listed,
		ListedPrice:   bitmap.ListedPrice,
		BitmapId:      bitmap.BitmapId,
		Content:       bitmap.Content,
		Source:        bitmap.Source,
		SourceUrl:     SOURCE_URL[bitmap.Source] + bitmap.InscriptionId,
	}
}

type BitmapListInfoResponseList struct {
	Count int64                    `json:"count"`
	Items []BitmapListInfoResponse `json:"items"`
}

func BuildBitmapListInfoResponseList(cnt int64, bitmaps []model.BitmapListInfo) BitmapListInfoResponseList {
	var list []BitmapListInfoResponse
	for _, bitmap := range bitmaps {
		list = append(list, BuildBitmapListInfoResponse(bitmap))
	}
	return BitmapListInfoResponseList{
		Count: cnt,
		Items: list,
	}
}
