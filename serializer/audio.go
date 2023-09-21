package serializer

import "server-template/model"

type AudioResponse struct {
	Id            int64  `json:"id"`
	InscriptionId string `json:"inscription_id"`
	ContentType   string `json:"content_type"`
}

func BuildAudioResponse(au model.Inscription) AudioResponse {
	return AudioResponse{
		Id:            au.Id,
		InscriptionId: au.InscriptionId,
		ContentType:   au.ContentType,
	}
}

type AudioListResponse struct {
	Count int
	Items []AudioResponse
}

func BuildAudioListResponse(audios []model.Inscription) AudioListResponse {
	var audioList []AudioResponse
	for _, audio := range audios {
		au := BuildAudioResponse(audio)
		audioList = append(audioList, au)
	}

	return AudioListResponse{
		Count: len(audioList),
		Items: audioList,
	}
}
