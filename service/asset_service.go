package service

import (
	"fmt"
	"server-template/config"
	"server-template/model"
	"server-template/serializer"
)

type AssetsListService struct {
	Address  string `json:"address" form:"address" binding:"required"`
	Type     string `json:"type" form:"type"`
	Category string `json:"category" form:"category"`
	Offset   int    `form:"offset,default=0" json:"offset"`
	Limit    int    `form:"limit,default=1" json:"limit"`
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
