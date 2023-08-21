package serializer

import (
	"server-template/model"
)

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

func BuildBitmapListResponse(items []model.Bitmap) []*BitmapResponse {
	var inscriptions []*BitmapResponse
	for _, item := range items {
		inscription := BuildBitmapResponse(item)
		inscriptions = append(inscriptions, inscription)
	}

	return inscriptions
}
