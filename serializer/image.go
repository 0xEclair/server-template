package serializer

import "server-template/model"

type ImageListResponse struct {
	Count int64                  `json:"count"`
	Items []*InscriptionResponse `json:"items"`
}

func BuildImageListResponse(cnt int64, items []model.Inscription) *ImageListResponse {
	return &ImageListResponse{
		Count: cnt,
		Items: BuildInscriptionListResponse(items),
	}
}

func BuildAddressInscriptionsListResponse(cnt int64, items []model.Inscription) *ImageListResponse {
	return &ImageListResponse{
		Count: cnt,
		Items: BuildInscriptionListResponse(items),
	}
}
