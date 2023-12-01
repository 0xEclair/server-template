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
