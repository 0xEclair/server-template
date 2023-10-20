package service

import (
	"server-template/config"
	"server-template/model"
	"server-template/serializer"
	"server-template/third/ordyssey"
	"sync"
	"time"
)

type AllBitmapService struct {
	Offset int `json:"offset" form:"offset,default=0"`
	Limit  int `json:"limit" form:"limit,default=1" `
}

var cacheBitmaps serializer.OrdysseyBitmapListResponse
var interval = 150 * time.Second
var last time.Time
var mu sync.Mutex

func (s AllBitmapService) List() serializer.Response {
	mu.Lock()
	if cacheBitmaps.Items != nil {
		if time.Now().Sub(last) < interval {
			start, end := s.calculateRange()
			cache := cacheBitmaps.Items[start:end]
			mu.Unlock()
			return serializer.Response{
				Code: 200,
				Data: serializer.OrdysseyBitmapListResponse{
					Cnt:   cacheBitmaps.Cnt,
					Items: cache,
				},
			}
		}
	}
	mu.Unlock()

	bitmaps, err := ordyssey.AllBitmaps()
	if err != nil {
		return serializer.Response{
			Code: 401,
			Err:  err.Error(),
		}
	}

	inscriptionIds := []string{}
	bitmapsFromOrdyssey := make(map[string]ordyssey.Bitmap)
	for _, bitmap := range bitmaps {
		inscriptionIds = append(inscriptionIds, bitmap.InscriptionId)
		if bitmapsFromOrdyssey[bitmap.InscriptionId].Price == 0 {
			bitmapsFromOrdyssey[bitmap.InscriptionId] = bitmap
		} else {
			if bitmapsFromOrdyssey[bitmap.InscriptionId].Price > bitmap.Price {
				bitmapsFromOrdyssey[bitmap.InscriptionId] = bitmap
			}
		}
	}

	var bms []model.Bitmap
	config.Postgres.Table("bitmap_holder").Select("id, bitmap_id, inscription_id").Where("inscription_id in ?", inscriptionIds).Order("bitmap_id").Find(&bms)

	mu.Lock()
	cacheBitmaps = serializer.BuildOrdysseyBitmapListResponse(bms, bitmapsFromOrdyssey)
	last = time.Now()
	mu.Unlock()

	start, end := s.calculateRange()
	cache := cacheBitmaps.Items[start:end]

	return serializer.Response{
		Code: 200,
		Data: serializer.OrdysseyBitmapListResponse{
			Cnt:   cacheBitmaps.Cnt,
			Items: cache,
		},
	}
}

func (s AllBitmapService) calculateRange() (int, int) {
	var start, end int
	if s.Offset > len(cacheBitmaps.Items) {
		return 0, 0
	}

	end = s.Offset + s.Limit
	if end > len(cacheBitmaps.Items) {
		end = len(cacheBitmaps.Items)
	}
	start = s.Offset
	return start, end
}
