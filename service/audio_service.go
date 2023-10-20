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
	Offset      int    `form:"offset,default=0" json:"offset"`
	Limit       int    `form:"limit,default=1" json:"limit"`
	ContentType string `form:"content_type" json:"content_type"`
}

func (s AudioService) List() serializer.Response {
	var realType []string
	if s.ContentType == "" {
		realType = []string{
			"audio/ogg",
			"audio/wav",
			"audio/flac",
			"audio/midi",
			"audio/mpeg",
			"audio/x-m4a",
			"audio/mod",
		}
	} else if s.ContentType == "model" {
		realType = []string{
			"model/gltf-binary",
		}
	}

	var audioInscriptions []model.Inscription
	config.Postgres.Where("id > 0 and address = ? and content_type in ? and oss_url is not null", s.Address, realType).Find(&audioInscriptions)

	var audioBRC420Entries []model.BRC420Entry
	config.Postgres.Where("address = ? and content_type in ?", s.Address, realType).Find(&audioBRC420Entries)

	var audioBRC420Inscriptions []model.Inscription
	set := make(map[int64]bool)
	if len(audioBRC420Entries) != 0 {
		var ids []string
		for _, abi := range audioBRC420Entries {
			ids = append(ids, abi.Ref)
			set[abi.Id] = true
		}
		config.Postgres.Where("inscription_id in ? and oss_url is not null", ids).Find(&audioBRC420Inscriptions)
	}
	// config.Postgres.
	// 	Table("brc420_entries").
	// 	Select("brc420_entries.id, brc420_entries.inscription_id, brc420_entries.address, brc420_entries.content_type, brc420_details.tick").
	// 	Joins("left join brc420_details on brc420_entries.ref=brc420_details.inscription_id").
	// 	Where("brc420_entries.address = ? and brc420_entries.content_type in ?", s.Address, audioType).Order("brc420_entries.id").Find(&audioBRC420Inscriptions)

	var realAudioInscriptions []model.Inscription
	for _, audio := range audioInscriptions {
		if set[audio.Id] {
			continue
		}

		realAudioInscriptions = append(realAudioInscriptions, audio)
	}
	realAudioInscriptions = append(realAudioInscriptions, audioBRC420Inscriptions...)
	return serializer.Response{
		Code: 200,
		Data: serializer.BuildAudioListResponse(realAudioInscriptions, s.Offset, s.Limit),
	}
}

func (s AudioService) ListModel() serializer.Response {
	var realType []string
	if s.ContentType == "" {
		realType = []string{
			"audio/ogg",
			"audio/wav",
			"audio/flac",
			"audio/midi",
			"audio/mpeg",
			"audio/x-m4a",
			"audio/mod",
		}
	} else if s.ContentType == "model" {
		realType = []string{
			"model/gltf-binary",
		}
	}

	// all music
	var audioInscriptions []model.Inscription
	config.Postgres.Where("id > 0 and address = ? and content_type in ? and oss_url is not null", s.Address, realType).Find(&audioInscriptions)

	// all brc420 music
	var audioBRC420Entries []model.BRC420EntryWithName
	config.Postgres.Table("brc420_entries").
		Select("brc420_entries.id", "brc420_entries.inscription_id", "brc420_entries.address", "brc420_entries.content_type", "brc420_entries.ref", "brc420_entries.fee_at", "brc420_entries.mint_timestamp", "brc420_details.name").
		Joins("left join brc420_details on brc420_entries.ref=brc420_details.tick").Where("brc420_entries.address = ? and content_type in ?", s.Address, realType).Find(&audioBRC420Entries)

	var audioBRC420Inscriptions []model.Inscription
	set := make(map[int64]bool)
	if len(audioBRC420Entries) != 0 {
		// brc420 map id true
		var ids []string
		for _, abi := range audioBRC420Entries {
			ids = append(ids, abi.Ref)
			set[abi.Id] = true
		}

		config.Postgres.Where("inscription_id in ? and oss_url is not null", ids).Find(&audioBRC420Inscriptions)
	}

	id2count := make(map[string]int64)
	id2name := make(map[string]string)
	for _, t := range audioBRC420Entries {
		id2count[t.Ref] += 1
		id2name[t.Ref] = t.Name
	}

	var realAudioInscriptions []model.InscriptionWithNameCount
	for _, audio := range audioInscriptions {
		if set[audio.Id] {
			continue
		}

		realAudioInscriptions = append(realAudioInscriptions, model.InscriptionWithNameCount{
			Id:            audio.Id,
			InscriptionId: audio.InscriptionId,
			Address:       audio.Address,
			Content:       audio.Content,
			ContentType:   audio.ContentType,
			OssUrl:        audio.OssUrl,
		})
	}

	for _, audio := range audioBRC420Inscriptions {
		name := id2name[audio.InscriptionId]
		count := id2count[audio.InscriptionId]
		realAudioInscriptions = append(realAudioInscriptions, model.InscriptionWithNameCount{
			Id:            audio.Id,
			InscriptionId: audio.InscriptionId,
			Address:       audio.Address,
			Content:       audio.Content,
			ContentType:   audio.ContentType,
			OssUrl:        audio.OssUrl,
			Name:          name,
			Count:         count,
		})
	}
	return serializer.Response{
		Code: 200,
		Data: serializer.BuildAudioListResponseWithName(realAudioInscriptions, s.Offset, s.Limit),
	}
}

func (s AudioService) CanMerge() serializer.Response {
	var realType []string
	if s.ContentType == "" {
		realType = []string{
			"audio/ogg",
			"audio/wav",
			"audio/flac",
			"audio/midi",
			"audio/mpeg",
			"audio/x-m4a",
			"audio/mod",
		}
	} else if s.ContentType == "model" {
		realType = []string{
			"model/gltf-binary",
		}
	}

	var audioInscriptions []model.Inscription
	config.Postgres.Where("id > 0 and address = ? and content_type in ? and oss_url is not null", s.Address, realType).Find(&audioInscriptions)

	var audioBRC420Entries []model.BRC420Entry
	config.Postgres.Where("address = ? and content_type in ?", s.Address, realType).Find(&audioBRC420Entries)

	var audioBRC420Inscriptions []model.Inscription
	set := make(map[int64]bool)
	if len(audioBRC420Entries) != 0 {
		var ids []string
		for _, abi := range audioBRC420Entries {
			ids = append(ids, abi.Ref)
			set[abi.Id] = true
		}
		config.Postgres.Where("inscription_id in ? and oss_url is not null", ids).Find(&audioBRC420Inscriptions)
	}
	// config.Postgres.
	// 	Table("brc420_entries").
	// 	Select("brc420_entries.id, brc420_entries.inscription_id, brc420_entries.address, brc420_entries.content_type, brc420_details.tick").
	// 	Joins("left join brc420_details on brc420_entries.ref=brc420_details.inscription_id").
	// 	Where("brc420_entries.address = ? and brc420_entries.content_type in ?", s.Address, audioType).Order("brc420_entries.id").Find(&audioBRC420Inscriptions)

	var realAudioInscriptions []model.Inscription
	for _, audio := range audioInscriptions {
		if set[audio.Id] {
			continue
		}

		realAudioInscriptions = append(realAudioInscriptions, audio)
	}
	realAudioInscriptions = append(realAudioInscriptions, audioBRC420Inscriptions...)
	res := false
	for _, audio := range realAudioInscriptions {
		if audio.InscriptionId == "acd5f7ae888e3ea2899533bcd0e91c3b74ca4e27d40ad9a01020ad1c6aa4b5adi0" {
			res = true
		}
	}

	return serializer.Response{
		Code: 200,
		Data: res,
	}
}
