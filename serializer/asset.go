package serializer

import "server-template/model"

type AssetResponse struct {
	Id            int64  `json:"id"`
	InscriptionId string `json:"inscription_id"`
	Address       string `json:"address"`
	Type          string `json:"type"`
	Category      string `json:"category"`
	Tag           string `json:"tag"`
	Collection    string `json:"collection"`
}

func BuildAssetResponse(inscription model.Asset) *AssetResponse {
	return &AssetResponse{
		Id:            inscription.Id,
		InscriptionId: inscription.InscriptionId,
		Address:       inscription.Address,
		Type:          inscription.Type,
		Category:      inscription.Category,
		Collection:    inscription.Collection,
		Tag:           inscription.Tag,
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

type AssetDLCResponse struct {
	Id            int64  `json:"id"`
	InscriptionId string `json:"inscription_id"`
	Address       string `json:"address"`
	Type          string `json:"type"`
	Category      string `json:"category"`
	Tag           string `json:"tag"`
	Collection    string `json:"collection"`
	Dlc           bool   `json:"dlc"`
}

type AssetsDLCListResponse struct {
	Count int64               `json:"count"`
	Items []*AssetDLCResponse `json:"items"`
}

func BuildAssetDLCResponse(inscription model.AssDLC) *AssetDLCResponse {
	return &AssetDLCResponse{
		Id:            inscription.Id,
		InscriptionId: inscription.InscriptionId,
		Address:       inscription.Address,
		Type:          inscription.Type,
		Category:      inscription.Category,
		Collection:    inscription.Collection,
		Tag:           inscription.Tag,
		Dlc:           inscription.DLC,
	}
}

func BuildAssetsListWithCntResponseDLC(cnt int64, items []model.AssDLC) *AssetsDLCListResponse {
	var inscriptions []*AssetDLCResponse
	for _, item := range items {
		inscription := BuildAssetDLCResponse(item)
		inscriptions = append(inscriptions, inscription)
	}

	return &AssetsDLCListResponse{
		Count: cnt,
		Items: inscriptions,
	}
}
