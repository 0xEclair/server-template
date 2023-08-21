package service

import (
	"server-template/config"
	"server-template/model"
	"server-template/serializer"
)

type BitmapService struct {
	BitmapIds []int32 `json:"id" form:"id" binding:"required"`
}

func (s *BitmapService) Find() serializer.Response {
	var bitmaps []model.Bitmap
	config.Postgres.Table("bitmap_holder").Where("bitmap_id in ?", s.BitmapIds).Find(&bitmaps)

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildBitmapListResponse(bitmaps),
	}
}

type BitmapCountService struct{}

func (s *BitmapCountService) Count() serializer.Response {
	var bitmaps model.Bitmap
	config.Postgres.Table("bitmap_holder").Order("bitmap_id desc").Limit(1).Find(&bitmaps)
	return serializer.Response{
		Code: 200,
		Data: bitmaps.BitmapId,
	}
}

type BitmapListService struct {
	Address string `json:"address" form:"address" binding:"required"`
}

func (s *BitmapListService) List() serializer.Response {
	var bitmaps []model.Bitmap
	config.Postgres.Table("bitmap_holder").Where("address = ?", s.Address).Find(&bitmaps)

	var cnt int64
	config.Postgres.Table("bitmap_holder").Where("address = ?", s.Address).Count(&cnt)
	return serializer.Response{
		Code: 200,
		Data: serializer.BuildBitmapListResponse(bitmaps),
	}
}

type BitmapRankService struct {
	Offset int `form:"offset,default=0" json:"offset"`
	Limit  int `form:"limit,default=20" json:"limit"`
}

func (s *BitmapRankService) List() serializer.Response {
	var rank []model.BitmapRank
	config.Postgres.Table("bitmap_holder").
		Select("address, count(address) as number").
		Group("address").
		Order("count(address) desc").
		Offset(s.Offset).
		Limit(s.Limit).
		Find(&rank)

	return serializer.Response{
		Code: 200,
		Data: rank,
	}
}
