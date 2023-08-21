package service

import (
	"server-template/config"
	"server-template/model"
	"server-template/serializer"
)

type ListInscriptionsService struct {
	Address string `form:"address" json:"address"`
	Offset  int    `form:"offset,default=0" json:"offset"`
	Limit   int    `form:"limit,default=20" json:"limit"`
}

func (s *ListInscriptionsService) List() serializer.Response {
	var inscriptions []model.Inscription
	config.Postgres.Where("address = ?", s.Address).Order("id asc").Offset(s.Offset).Limit(s.Limit).Find(&inscriptions)

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildInscriptionListResponse(inscriptions),
	}
}
