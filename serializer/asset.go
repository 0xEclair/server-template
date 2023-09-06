package serializer

import "server-template/model"

type AssetResponse struct {
	Id            int64  `json:"id"`
	InscriptionId string `json:"inscription_id"`
	Address       string `json:"address"`
	Type          string `json:"type"`
	Category      string `json:"category"`
}

func BuildAssetResponse(inscription model.Asset) *AssetResponse {
	return &AssetResponse{
		Id:            inscription.Id,
		InscriptionId: inscription.InscriptionId,
		Address:       inscription.Address,
		Type:          inscription.Type,
		Category:      inscription.Category,
	}
}

type AssetsListResponse struct {
	Count int64            `json:"count"`
	Items []*AssetResponse `json:"items"`
}

func BuildAssetsListWithCntResponse(cnt int64, items []model.Asset) *AssetsListResponse {
	var inscriptions []*AssetResponse
	for _, item := range items {
		inscription := BuildAssetResponse(item)
		inscriptions = append(inscriptions, inscription)
	}

	return &AssetsListResponse{
		Count: cnt,
		Items: inscriptions,
	}
}
