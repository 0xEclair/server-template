package serializer

import (
	"server-template/model"
	"sort"
)

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

func BuildAudioListResponse(audios []model.Inscription, offset, limit int) AudioListResponse {
	set := make(map[int64]bool)
	var audioList []AudioResponse
	for _, audio := range audios {
		if set[audio.Id] {
			continue
		}
		au := BuildAudioResponse(audio)
		audioList = append(audioList, au)
		set[audio.Id] = true
	}

	sort.Slice(audioList, func(i, j int) bool {
		return audioList[i].Id < audioList[j].Id
	})

	var last int
	if offset+limit > len(audioList) {
		last = len(audioList)
	} else {
		last = offset + limit
	}

	return AudioListResponse{
		Count: len(audioList),
		Items: audioList[offset:last],
	}
}
