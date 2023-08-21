package serializer

import "server-template/model"

type AvatarListResponse struct {
	Count int64                  `json:"count"`
	Items []*InscriptionResponse `json:"items"`
}

func BuildAvatarListResponse(cnt int64, items []model.Inscription) *ImageListResponse {
	return &ImageListResponse{
		Count: cnt,
		Items: BuildInscriptionListResponse(items),
	}
}
