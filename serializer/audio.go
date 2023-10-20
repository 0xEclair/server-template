package serializer

import (
	"server-template/model"
	"sort"
)

type AudioResponse struct {
	Id            int64  `json:"id"`
	InscriptionId string `json:"inscription_id"`
	ContentType   string `json:"content_type"`
	Name          string `json:"name,omitempty"`
	Count         int64  `json:"count,omitempty"`
}

func BuildAudioResponse(au model.Inscription) AudioResponse {
	return AudioResponse{
		Id:            au.Id,
		InscriptionId: au.InscriptionId,
		ContentType:   au.ContentType,
	}
}

func BuildAudioResponseWithName(au model.InscriptionWithNameCount) AudioResponse {
	return AudioResponse{
		Id:            au.Id,
		InscriptionId: au.InscriptionId,
		ContentType:   au.ContentType,
		Name:          au.Name,
		Count:         au.Count,
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

	lenAudioList := len(audioList)
	last := lenAudioList

	items := make([]AudioResponse, 0)
	if offset+limit < lenAudioList {
		last = offset + limit
	}

	if offset < lenAudioList {
		items = audioList[offset:last]
	}

	return AudioListResponse{
		Count: lenAudioList,
		Items: items,
	}
}

func BuildAudioListResponseWithName(audios []model.InscriptionWithNameCount, offset, limit int) AudioListResponse {
	//	set := make(map[int64]bool)
	var audioList []AudioResponse
	for _, audio := range audios {
		// if set[audio.Id] {
		// 	continue
		// }
		au := BuildAudioResponseWithName(audio)
		audioList = append(audioList, au)
		//set[audio.Id] = true
	}

	sort.Slice(audioList, func(i, j int) bool {
		return audioList[i].Id < audioList[j].Id
	})

	lenAudioList := len(audioList)
	last := lenAudioList

	items := make([]AudioResponse, 0)
	if offset+limit < lenAudioList {
		last = offset + limit
	}

	if offset < lenAudioList {
		items = audioList[offset:last]
	}

	return AudioListResponse{
		Count: lenAudioList,
		Items: items,
	}
}
