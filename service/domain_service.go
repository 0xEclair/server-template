package service

import (
	"strconv"
	"strings"

	"server-template/config"
	"server-template/model"
	"server-template/serializer"
	"server-template/third/bestinslot"
)

type BRC20 struct {
	P    string `json:"p"`
	Op   string `json:"op"`
	Tick string `json:"tick`
	Amt  string `json:"amt"`
}

type SATS struct {
	P    string `json:"p"`
	Op   string `json:"op"`
	Name string `json:"name"`
}

type DomainService struct {
	Offset  int    `form:"offset,default=0" json:"offset"`
	Limit   int    `form:"limit,default=20" json:"limit"`
	Address string `form:"address" json:"address"`
}

func (s *DomainService) List() serializer.Response {
	var tempInscriptions []model.Inscription
	config.Postgres.Select("id", "inscription_id", "content").Where("address = ? and content_type in ?", s.Address, []string{"text/plain", "text/plain;charset=utf-8"}).Order("id asc").Find(&tempInscriptions)

	var verifiedInscriptions []model.Inscription
	var unverifiedInscriptions []model.Inscription
	for _, inscription := range tempInscriptions {
		// var sats SATS
		// err := json.Unmarshal([]byte(inscription.Content), &sats)
		// if err != nil {

		// } else {
		// 	if sats.P == "sns" && sats.Op == "reg" && strings.HasSuffix(sats.Name, ".sats") {
		// 		var insc model.Inscription
		// 		config.Postgres.Select("id").Where("content in ?", []string{inscription.Content, sats.Name}).First(&insc)
		// 		inscription.Content = sats.Name
		// 		if inscription.Id == insc.Id {
		// 			verifiedInscriptions = append(verifiedInscriptions, inscription)
		// 		} else {
		// 			unverifiedInscriptions = append(unverifiedInscriptions, inscription)
		// 		}

		// 		continue
		// 	}
		// }

		if strings.HasSuffix(inscription.Content, ".sats") {
			// var insc model.Inscription
			// config.Postgres.Select("id").Where("content = ?", inscription.Content).First(&insc)
			// if inscription.Id == insc.Id {
			// 	verifiedInscriptions = append(verifiedInscriptions, inscription)
			// } else {
			// 	unverifiedInscriptions = append(unverifiedInscriptions, inscription)
			// }

			// skip .sats
			continue
		}

		if strings.HasSuffix(inscription.Content, ".btc") {
			var insc model.Inscription
			config.Postgres.Select("id").Where("content = ?", inscription.Content).First(&insc)
			if inscription.Id == insc.Id {
				verifiedInscriptions = append(verifiedInscriptions, inscription)
			} else {
				unverifiedInscriptions = append(unverifiedInscriptions, inscription)
			}

			continue
		}

		if strings.HasSuffix(inscription.Content, ".xbt") {
			var insc model.Inscription
			config.Postgres.Select("id").Where("content = ?", inscription.Content).First(&insc)
			if inscription.Id == insc.Id {
				verifiedInscriptions = append(verifiedInscriptions, inscription)
			} else {
				unverifiedInscriptions = append(unverifiedInscriptions, inscription)
			}

			continue
		}

		if strings.HasSuffix(inscription.Content, ".gm") {
			var insc model.Inscription
			config.Postgres.Select("id").Where("content = ?", inscription.Content).First(&insc)
			if inscription.Id == insc.Id {
				verifiedInscriptions = append(verifiedInscriptions, inscription)
			} else {
				unverifiedInscriptions = append(unverifiedInscriptions, inscription)
			}

			continue
		}

		if strings.HasSuffix(inscription.Content, ".bitter") {
			var insc model.Inscription
			config.Postgres.Select("id").Where("content = ?", inscription.Content).First(&insc)
			if inscription.Id == insc.Id {
				verifiedInscriptions = append(verifiedInscriptions, inscription)
			} else {
				unverifiedInscriptions = append(unverifiedInscriptions, inscription)
			}
		}

		if strings.HasSuffix(inscription.Content, ".x") {
			var insc model.Inscription
			config.Postgres.Select("id").Where("content = ?", inscription.Content).First(&insc)
			if inscription.Id == insc.Id {
				verifiedInscriptions = append(verifiedInscriptions, inscription)
			} else {
				unverifiedInscriptions = append(unverifiedInscriptions, inscription)
			}

			continue
		}

		if strings.HasSuffix(inscription.Content, ".magic") {
			var insc model.Inscription
			config.Postgres.Select("id").Where("content = ?", inscription.Content).First(&insc)
			if inscription.Id == insc.Id {
				verifiedInscriptions = append(verifiedInscriptions, inscription)
			} else {
				unverifiedInscriptions = append(unverifiedInscriptions, inscription)
			}

			continue
		}

		if strings.HasSuffix(inscription.Content, ".ord") {
			var insc model.Inscription
			config.Postgres.Select("id").Where("content = ?", inscription.Content).First(&insc)
			if inscription.Id == insc.Id {
				verifiedInscriptions = append(verifiedInscriptions, inscription)
			} else {
				unverifiedInscriptions = append(unverifiedInscriptions, inscription)
			}

			continue
		}

		if strings.HasSuffix(inscription.Content, ".pokemon") {
			domains := strings.Split(inscription.Content, ".")
			i, err := strconv.ParseInt(domains[0], 10, 32)
			if err != nil {
				continue
			}
			if i < 1 || i > 1010 {
				continue
			}
			var insc model.Inscription
			config.Postgres.Select("id").Where("content = ? and id > 0", inscription.Content).First(&insc)
			if inscription.Id == insc.Id {
				verifiedInscriptions = append(verifiedInscriptions, inscription)
			} else {
				unverifiedInscriptions = append(unverifiedInscriptions, inscription)
			}

			continue
		}
	}

	sats, err := bestinslot.WalletSats(s.Address)
	if err != nil {
		// todo log
	}

	for _, satsname := range sats {
		inscription := model.Inscription{
			Id:            satsname.Id,
			InscriptionId: satsname.InscriptionId,
			Content:       satsname.Name,
		}

		verifiedInscriptions = append(verifiedInscriptions, inscription)
	}

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildDomainListResponse(verifiedInscriptions, unverifiedInscriptions),
	}
}
