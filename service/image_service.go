package service

import (
	"server-template/config"
	"server-template/model"
	"server-template/serializer"
)

type ImageService struct {
	Offset int    `form:"offset" json:"offset"`
	Limit  int    `form:"limit" json:"limit"`
	Type   string `form:"type" json:"type"`
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
		"image/svg+xml;charset=utf-8",
	}

	// config.Postgres.Select("id", "inscription_id", "address", "content", "content_type").
	// 	Where("id >= ? and content_type in ?", 0, typeList).
	// 	Offset(s.Offset).Order("id asc").
	// 	Limit(s.Limit).
	// 	Find(&inscriptions)

	var cnt int64
	// config.Postgres.Model(&model.Inscription{}).
	// 	Where("id >= ? and content_type in ?", 0, typeList).
	// 	Count(&cnt)
	_ = typeList
	_ = cnt

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildImageListResponse(1204385, inscriptions),
	}
}

func (s *ImageService) ListWithHTML() serializer.Response {
	var inscriptions []model.Inscription
	typeList := []string{
		"image/avif",
		"image/gif",
		"image/jpeg",
		"image/png",
		"image/svg+xml",
		"image/webp",
		"image/svg+xml;charset=utf-8",
		"text/html",
		"text/html;charset=utf-8",
	}

	config.Postgres.Select("id", "inscription_id", "address", "content", "content_type").
		Where("id >= ? and content_type in ?", 0, typeList).
		Offset(s.Offset).Order("id asc").
		Limit(s.Limit).
		Find(&inscriptions)

	var cnt int64
	config.Postgres.Model(&model.Inscription{}).
		Where("id >= ? and content_type in ?", 0, typeList).
		Count(&cnt)

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildImageListResponse(cnt, inscriptions),
	}
}
