package service

import (
	"server-template/cache"
	"server-template/config"
	"server-template/model"
	"server-template/serializer"
)

type AssetsCollectionService struct {
	Address string `json:"address" form:"address" binding:"required"`
}

func (s *AssetsCollectionService) Collection() serializer.Response {
	dlcList := []string{}
	for k, _ := range cache.DLCToAssets {
		if k == "dbdbc1ff9fa94d1149c240f742fa444a853b0101d3fb898376e00adb8792454di0" {
			dlcList = append(dlcList, "e80eadea13e2175949168c279f7a47a467b6c5e00a84d45dd8ae40aefda89fe5i0")
		} else {
			dlcList = append(dlcList, k)
		}
	}

	var brc420Entries []model.BRC420Entry
	config.Postgres.Where("address = ? and ref in ?", s.Address, dlcList).Find(&brc420Entries)

	refList := []string{}
	for _, entry := range brc420Entries {
		refList = append(refList, entry.Ref)
	}

	var brc420Details []model.BRC420Detail
	config.Postgres.Where("tick in ?", refList).Find(&brc420Details)

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildAssetCollectionListResponse(brc420Details),
	}
}

func (s *AssetsCollectionService) CollectionV2() serializer.Response {
	dlcList := []string{}
	for k, _ := range cache.DLCToAssets {
		if k == "dbdbc1ff9fa94d1149c240f742fa444a853b0101d3fb898376e00adb8792454di0" {
			dlcList = append(dlcList, "e80eadea13e2175949168c279f7a47a467b6c5e00a84d45dd8ae40aefda89fe5i0")
		} else {
			dlcList = append(dlcList, k)
		}
	}

	var brc420Entries []model.BRC420Entry
	config.Postgres.Where("address = ? and ref in ?", s.Address, dlcList).Find(&brc420Entries)

	refList := []string{}
	for _, entry := range brc420Entries {
		refList = append(refList, entry.Ref)
	}

	var brc420Details []model.BRC420Detail
	config.Postgres.Where("tick in ?", refList).Find(&brc420Details)

	nameToInsids := make(map[string][]string)
	{
		var assets []model.Asset
		config.Postgres.Model(&model.Asset{}).Where("address = ? and cons = true", s.Address).Order("id").Find(&assets)

		list := []string{}
		for _, asset := range assets {
			list = append(list, asset.InscriptionId)
		}

		var brc420Details []model.BRC420Detail
		config.Postgres.Model(&model.BRC420Detail{}).Where("address = ? and tick in ?", s.Address, list).Find(&brc420Details)

		for _, detail := range brc420Details {
			nameToInsids[detail.Name] = append(nameToInsids[detail.Name], detail.Tick)
		}
	}

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildAssetCollectionNameListResponse(brc420Details, nameToInsids),
	}
}

type CollectionDLCService struct {
	InscriptionId string `json:"inscription_id" form:"inscription_id"`
	Offset        int    `json:"offset" form:"offset,default=0"`
	Limit         int    `json:"limit" form:"limit,default=1"`
}

func (s *CollectionDLCService) List() serializer.Response {
	assets := cache.DLCToAssets[s.InscriptionId]

	var res []model.Asset
	for _, asset := range assets {
		if asset.Cons {
			res = append(res, asset)
		}
	}

	cnt := len(res)
	last := s.Offset + s.Limit
	if last > len(res) {
		last = len(res)
	}

	if s.Offset < len(res) {
		res = res[s.Offset:last]
	} else {
		res = make([]model.Asset, 0)
	}

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildAssetsListWithCntResponse(int64(cnt), res),
	}
}
