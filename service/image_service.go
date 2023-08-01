package service

import (
	"server-template/config"
	"server-template/model"
	"server-template/serializer"
)

type ImageService struct {
	Offset int `form:"offset" json:"offset"`
	Limit  int `form:"limit" json:"limit"`
}

func (s *ImageService) List() serializer.Response {
	var inscriptions []model.Inscription
	typeList := []string{
		"image/avif",
		"image/gif",
		"image/jpeg",
		"image/png",
		"image/svg+xml",
		"image/webp",
	}
	config.Postgres.Where("content_type in ? and id > ?", typeList, 0).
		Offset(s.Offset).
		Limit(s.Limit).
		Find(&inscriptions)

	return serializer.Response{
		Code: 200,
		Data: inscriptions,
	}
}
