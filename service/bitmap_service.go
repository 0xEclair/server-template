package service

type BipmapService struct {
	Number string `form:"number" json:"number"`
	Id     string `form:"id" json:"id"`
}

// func (s *InscriptionService) Find() serializer.Response {
// 	var inscription model.Inscription
// 	if s.Number != "" {
// 		id, err := strconv.ParseInt(s.Number, 10, 64)
// 		if err != nil {
// 			return serializer.Response{
// 				Code: 401,
// 				Err:  err.Error(),
// 			}
// 		}
// 		config.Postgres.Where("id = ?", id).First(&inscription, id)
// 	} else {
// 		config.Postgres.Where("inscription_id = ?", s.Id).First(&inscription)
// 	}

// 	return serializer.Response{
// 		Code: 200,
// 		Data: serializer.BuildInscriptionResponse(inscription),
// 	}
// }
