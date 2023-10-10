package service

import (
	"server-template/config"
	"server-template/model"
	"server-template/serializer"
	"server-template/third/ordyssey"
	"time"
)

type AllBitmapService struct{}

var cacheBitmaps serializer.OrdysseyBitmapListResponse
var interval = 6 * time.Minute
var last time.Time

func (s AllBitmapService) List() serializer.Response {
	if cacheBitmaps.Items != nil {
		if time.Now().Sub(last) < interval {
			return serializer.Response{
				Code: 200,
				Data: cacheBitmaps,
			}
		}
	}

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

	cacheBitmaps = serializer.BuildOrdysseyBitmapListResponse(bms, bitmapsFromOrdyssey)
	last = time.Now()
	return serializer.Response{
		Code: 200,
		Data: cacheBitmaps,
	}
}
