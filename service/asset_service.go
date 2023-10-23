package service

import (
	"fmt"
	"server-template/config"
	"server-template/model"
	"server-template/serializer"
	"sort"
)

type AssetsListService struct {
	Address    string `json:"address" form:"address" binding:"required"`
	Type       string `json:"type" form:"type"`
	Category   string `json:"category" form:"category"`
	Collection string `json:"collection" form:"collection"`
	Tag        string `json:"tag" form:"tag"`
	Offset     int    `form:"offset,default=0" json:"offset"`
	Limit      int    `form:"limit,default=1" json:"limit"`
}

func (s *AssetsListService) List() serializer.Response {
	var assets []model.Asset

	w := fmt.Sprintf("address = '%s' and id >= 0", s.Address)
	if s.Type != "" {
		w = fmt.Sprintf("%s and type = '%s'", w, s.Type)
	}

	if s.Category != "" {
		w = fmt.Sprintf("%s and category = '%s'", w, s.Category)
	}

	config.Postgres.Where(w).Order("id").Offset(s.Offset).Limit(s.Limit).Find(&assets)

	var cnt int64
	config.Postgres.Table("assets").Where(w).Count(&cnt)
	return serializer.Response{
		Code: 200,
		Data: serializer.BuildAssetsListWithCntResponse(cnt, assets),
	}
}

func (s *AssetsListService) ListWithOss() serializer.Response {
	var assets []model.Asset

	w := fmt.Sprintf("assets.address = '%s' and assets.id >= 0 and inscriptions.oss_url is not null", s.Address)
	if s.Type != "" {
		w = fmt.Sprintf("%s and type in ('%s', 'character')", w, s.Type)
	}

	if s.Category != "" {
		w = fmt.Sprintf("%s and category = '%s'", w, s.Category)
	}

	if s.Collection != "" {
		w = fmt.Sprintf("%s and collection = '%s'", w, s.Collection)
	}

	if s.Tag != "" {
		w = fmt.Sprintf("%s and tag like '%%%s%%'", w, s.Tag)
	}

	config.Postgres.Table("assets").Select("assets.id, assets.inscription_id, assets.address, assets.type, assets.category, assets.collection, assets.tag").Joins("left join inscriptions on assets.id=inscriptions.id").Where(w).Order("id").Offset(s.Offset).Limit(s.Limit).Find(&assets)

	var cnt int64
	config.Postgres.Table("assets").Select("assets.id, assets.inscription_id, assets.address, assets.type, assets.category, assets.collection, assets.tag").Joins("left join inscriptions on assets.id=inscriptions.id").Where(w).Count(&cnt)
	return serializer.Response{
		Code: 200,
		Data: serializer.BuildAssetsListWithCntResponse(cnt, assets),
	}
}

func (s *AssetsListService) ListWithOssAndBRC420() serializer.Response {
	var assets []model.Asset

	w := fmt.Sprintf("assets.address = '%s' and assets.id >= 0 and inscriptions.oss_url is not null", s.Address)
	if s.Type != "" {
		w = fmt.Sprintf("%s and type in ('%s', 'character')", w, s.Type)
	}

	if s.Category != "" {
		w = fmt.Sprintf("%s and category = '%s'", w, s.Category)
	}

	if s.Collection != "" {
		w = fmt.Sprintf("%s and collection = '%s'", w, s.Collection)
	}

	if s.Tag != "" {
		w = fmt.Sprintf("%s and tag like '%%%s%%'", w, s.Tag)
	}

	config.Postgres.Table("assets").Select("assets.id, assets.inscription_id, assets.address, assets.type, assets.category, assets.collection, assets.tag").Joins("left join inscriptions on assets.id=inscriptions.id").Where(w).Order("id").Offset(s.Offset).Limit(s.Limit).Find(&assets)

	var cnt int64
	config.Postgres.Table("assets").Select("assets.id, assets.inscription_id, assets.address, assets.type, assets.category, assets.collection, assets.tag").Joins("left join inscriptions on assets.id=inscriptions.id").Where(w).Count(&cnt)

	var brc420AssetsWithName []model.BRC420EntryWithName
	config.Postgres.Table("brc420_entries").
		Select("brc420_entries.id", "brc420_entries.inscription_id", "brc420_entries.address", "brc420_entries.content_type", "brc420_entries.ref", "brc420_entries.fee_at", "brc420_entries.mint_timestamp", "brc420_details.name").
		Joins("left join brc420_details on brc420_entries.ref=brc420_details.tick").Where("brc420_entries.address = ? and content_type in ?", s.Address, []string{"text/html", "text/html;charset=utf-8"}).Find(&brc420AssetsWithName)

	findAllAssets := []string{}
	for _, asset := range brc420AssetsWithName {
		findAllAssets = append(findAllAssets, asset.Ref)
	}
	var brc420Assets []model.Asset
	config.Postgres.Table("assets").Select("assets.id, assets.inscription_id, assets.address, assets.type, assets.category, assets.collection, assets.tag").Where("assets.inscription_id in ? and id > 0", findAllAssets).Order("id").Find(&brc420Assets)


	
	assets = append(assets, brc420Assets...)

	sort.Slice(assets, func(i, j int) bool {
		return assets[i].Id < assets[j].Id
	})
	
	return serializer.Response{
		Code: 200,
		Data: serializer.BuildAssetsListWithCntResponse(cnt, assets),
	}
}
