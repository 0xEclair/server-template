package service

import (
	"server-template/config"
	"server-template/model"
	"server-template/serializer"
)

type AvatarService struct {
	Offset  int    `form:"offset" json:"offset"`
	Limit   int    `form:"limit" json:"limit"`
	Address string `form:"address" json:"address"`
}

func (s *AvatarService) List() serializer.Response {
	var inscriptions []model.Inscription
	typeList := []string{
		"image/avif",
		"image/gif",
		"image/jpeg",
		"image/png",
		"image/svg+xml",
		"image/webp",
		"image/svg+xml;charset=utf-8",
	}

	config.Postgres.Select("id", "inscription_id", "address", "content_type").
		Where("address = ? and id >= 0 and content_type in ?", s.Address, typeList).
		Offset(s.Offset).Order("id asc").
		Limit(s.Limit).
		Find(&inscriptions)

	var cnt int64
	config.Postgres.Model(&model.Inscription{}).
		Where("address = ? and id >= 0 and content_type in ?", s.Address, typeList).
		Count(&cnt)

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildAvatarListResponse(cnt, inscriptions),
	}
}
