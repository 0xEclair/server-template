package service

import (
	"math/big"
	"regexp"
	"strconv"
	"strings"

	"server-template/config"
	"server-template/model"
	"server-template/serializer"
	"server-template/third/bestinslot"
)

type InscriptionService struct {
	Number string `form:"number" json:"number"`
	Id     string `form:"id" json:"id"`
}

func (s *InscriptionService) Find() serializer.Response {
	var inscription model.Inscription
	if s.Number != "" {
		id, err := strconv.ParseInt(s.Number, 10, 64)
		if err != nil {
			return serializer.Response{
				Code: 401,
				Err:  err.Error(),
			}
		}
		config.Postgres.Where("id = ?", id).First(&inscription, id)
	} else {
		config.Postgres.Where("inscription_id = ?", s.Id).First(&inscription)
	}

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildInscriptionResponse(inscription),
	}
}

type AddressByConditionService struct {
	Condition string `uri:"condition"`
}

func (s *AddressByConditionService) Find() serializer.Response {
	if regexp.MustCompile(`^(bc1|[13])[a-zA-Z0-9]*$`).MatchString(s.Condition) {
		var inscription model.Inscription
		inscription.Address = s.Condition
		return serializer.Response{
			Code: 200,
			Data: serializer.BuildInscriptionWithoutContentResponse(inscription),
		}
	}
	t := s.Check()
	cond := s.Condition
	var inscription model.Inscription

	if t == "id" {
		config.Postgres.Where("id = ?", cond).First(&inscription)
	} else if t == "content" {
		config.Postgres.Where("id >= 0 and content = ?", cond).Order("id asc").Limit(1).First(&inscription)
	} else if t == "sats" {
		addr, _ := bestinslot.AddressBySats(cond)
		inscription.Address = addr
	} else {
		config.Postgres.Where("inscription_id = ?", cond).First(&inscription)
	}

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildInscriptionWithoutContentResponse(inscription),
	}
}

func (s *AddressByConditionService) Check() string {
	cond := s.Condition
	_, ok := new(big.Int).SetString(cond, 10)
	if ok {
		return "id"
	}

	if strings.Contains(cond, ".") {
		domain := strings.Split(cond, ".")
		if domain[1] == "sats" {
			return "sats"
		}
		return "content"
	}

	return "inscription_id"
}
