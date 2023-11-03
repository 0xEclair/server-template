package serializer

import (
	"server-template/model"
)

type InscriptionResponse struct {
	Id            int64  `json:"id"`
	InscriptionId string `json:"inscription_id"`
	Address       string `json:"address"`
	Content       string `json:"content,omitempty"`
	ContentType   string `json:"content_type,omitempty"`
}

type VerifiedInscriptionResponse struct {
	Id            int64  `json:"id"`
	InscriptionId string `json:"inscription_id"`
	Address       string `json:"address"`
	Content       string `json:"content,omitempty"`
	ContentType   string `json:"content_type,omitempty"`
	Verified      bool   `json:"verified"`
}

func BuildInscriptionWithoutContentResponse(inscription model.Inscription) *InscriptionResponse {
	return &InscriptionResponse{
		Id:            inscription.Id,
		InscriptionId: inscription.InscriptionId,
		Address:       inscription.Address,
	}
}

func BuildInscriptionResponse(inscription model.Inscription) *InscriptionResponse {
	return &InscriptionResponse{
		Id:            inscription.Id,
		InscriptionId: inscription.InscriptionId,
		Address:       inscription.Address,
		Content:       inscription.Content,
		ContentType:   inscription.ContentType,
	}
}

func BuildInscriptionListResponse(items []model.Inscription) []*InscriptionResponse {
	var inscriptions []*InscriptionResponse
	for _, item := range items {
		inscription := BuildInscriptionResponse(item)
		inscriptions = append(inscriptions, inscription)
	}

	return inscriptions
}

func BuildInscriptionListResponseForBitmap(items []model.Inscription) []*InscriptionResponse {
	var inscriptions []*InscriptionResponse
	for _, item := range items {
		item.Content = ""
		inscription := BuildInscriptionResponse(item)
		inscriptions = append(inscriptions, inscription)
	}

	return inscriptions
}

func BuildDomainResponse(inscription model.Inscription, verified bool) *VerifiedInscriptionResponse {
	return &VerifiedInscriptionResponse{
		Id:            inscription.Id,
		InscriptionId: inscription.InscriptionId,
		Address:       inscription.Address,
		Content:       inscription.Content,
		ContentType:   inscription.ContentType,
		Verified:      verified,
	}
}

func BuildDomainListResponse(verified []model.Inscription, unverified []model.Inscription) []*VerifiedInscriptionResponse {
	var inscriptions []*VerifiedInscriptionResponse
	for _, item := range verified {
		inscription := BuildDomainResponse(item, true)
		inscriptions = append(inscriptions, inscription)
	}

	for _, item := range unverified {
		inscription := BuildDomainResponse(item, false)
		inscriptions = append(inscriptions, inscription)
	}

	return inscriptions
}

func BuildVerifiedDomainListResponse(verified []model.Inscription) []*VerifiedInscriptionResponse {
	var inscriptions []*VerifiedInscriptionResponse
	for _, item := range verified {
		inscription := BuildDomainResponse(item, true)
		inscriptions = append(inscriptions, inscription)
	}

	return inscriptions
}
