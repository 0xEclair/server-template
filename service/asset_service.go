package service

import (
	"fmt"
	"server-template/config"
	"server-template/model"
	"server-template/serializer"
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

	config.Postgres.Table("assets").Select("assets.id, assets.inscription_id, assets.address, assets.type, assets.category").Joins("left join inscriptions on assets.id=inscriptions.id").Where(w).Order("id").Offset(s.Offset).Limit(s.Limit).Find(&assets)

	var cnt int64
	config.Postgres.Table("assets").Select("assets.id, assets.inscription_id, assets.address, assets.type, assets.category").Joins("left join inscriptions on assets.id=inscriptions.id").Where(w).Count(&cnt)
	return serializer.Response{
		Code: 200,
		Data: serializer.BuildAssetsListWithCntResponse(cnt, assets),
	}
}
