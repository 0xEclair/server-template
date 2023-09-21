package service

import (
	"server-template/config"
	"server-template/model"
	"server-template/serializer"
)

type AudioService struct {
	Address string `json:"address" form:"address" binding:"required"`
	//Type     string `json:"type" form:"type"`
	//Category string `json:"category" form:"category"`
	Offset int `form:"offset,default=0" json:"offset"`
	Limit  int `form:"limit,default=1" json:"limit"`
}

func (s AudioService) List() serializer.Response {
	audioType := []string{
		"audio/ogg",
		"audio/wav",
		"audio/flac",
		"audio/midi",
		"audio/mpeg",
		"audio/x-m4a",
		"audio/mod",
	}

	var audioInscriptions []model.Inscription
	config.Postgres.Where("id > 0 and address = ? and content_type in ?", s.Address, audioType).Order("id").Find(&audioInscriptions)

	//var audioBRC420Inscriptions []model.BRC420Entry
	// config.Postgres.
	// 	Table("brc420_entries").
	// 	Select("brc420_entries.id, brc420_entries.inscription_id, brc420_entries.address, brc420_entries.content_type, brc420_details.tick").
	// 	Joins("left join brc420_details on brc420_entries.ref=brc420_details.inscription_id").
	// 	Where("brc420_entries.address = ? and brc420_entries.content_type in ?", s.Address, audioType).Order("brc420_entries.id").Find(&audioBRC420Inscriptions)

	//config.Postgres.Where("address = ? and content_type in ?", s.Address, audioType).Order("id").Find(&audioBRC420Inscriptions)

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildAudioListResponse(audioInscriptions),
	}
}
