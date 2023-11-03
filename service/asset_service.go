package service

import (
	"fmt"
	"server-template/cache"
	"server-template/config"
	"server-template/model"
	"server-template/serializer"
	"sort"
)

type AssetsListService struct {
	Address    string `json:"address" form:"address" binding:"required"`
	Type       string `json:"type" form:"type"`
	Category   string `json:"category" form:"category"`
	Collection string `json:"collection" form:"collection"`
	Tag        string `json:"tag" form:"tag"`
	Offset     int    `form:"offset,default=0" json:"offset"`
	Limit      int    `form:"limit,default=1" json:"limit"`
}

func (s *AssetsListService) List() serializer.Response {
	var assets []model.Asset

	w := fmt.Sprintf("address = '%s' and id >= 0", s.Address)
	if s.Type != "" {
		w = fmt.Sprintf("%s and type = '%s'", w, s.Type)
	}

	if s.Category != "" {
		w = fmt.Sprintf("%s and category = '%s'", w, s.Category)
	}

	config.Postgres.Where(w).Order("id").Offset(s.Offset).Limit(s.Limit).Find(&assets)

	var cnt int64
	config.Postgres.Table("assets").Where(w).Count(&cnt)
	return serializer.Response{
		Code: 200,
		Data: serializer.BuildAssetsListWithCntResponse(cnt, assets),
	}
}

func (s *AssetsListService) ListWithOss() serializer.Response {
	var assets []model.Asset

	w := fmt.Sprintf("assets.address = '%s' and assets.id >= 0 and inscriptions.oss_url is not null", s.Address)
	if s.Type != "" {
		w = fmt.Sprintf("%s and type in ('%s', 'character')", w, s.Type)
	}

	if s.Category != "" {
		w = fmt.Sprintf("%s and category = '%s'", w, s.Category)
	}

	if s.Collection != "" {
		w = fmt.Sprintf("%s and collection = '%s'", w, s.Collection)
	}

	if s.Tag != "" {
		w = fmt.Sprintf("%s and tag like '%%%s%%'", w, s.Tag)
	}

	config.Postgres.Table("assets").Select("assets.id, assets.inscription_id, assets.address, assets.type, assets.category, assets.collection, assets.tag").Joins("left join inscriptions on assets.id=inscriptions.id").Where(w).Order("id").Offset(s.Offset).Limit(s.Limit).Find(&assets)

	var cnt int64
	config.Postgres.Table("assets").Select("assets.id, assets.inscription_id, assets.address, assets.type, assets.category, assets.collection, assets.tag").Joins("left join inscriptions on assets.id=inscriptions.id").Where(w).Count(&cnt)
	return serializer.Response{
		Code: 200,
		Data: serializer.BuildAssetsListWithCntResponse(cnt, assets),
	}
}

func (s *AssetsListService) ListWithOssAndBRC420() serializer.Response {
	var assets []model.Asset

	w := fmt.Sprintf("assets.address = '%s' and assets.id >= 0 and inscriptions.oss_url is not null", s.Address)
	if s.Type != "" {
		w = fmt.Sprintf("%s and type in ('%s', 'character')", w, s.Type)
	}

	if s.Category != "" {
		w = fmt.Sprintf("%s and category = '%s'", w, s.Category)
	}

	if s.Collection != "" {
		w = fmt.Sprintf("%s and collection = '%s'", w, s.Collection)
	}

	if s.Tag != "" {
		w = fmt.Sprintf("%s and tag like '%%%s%%'", w, s.Tag)
	}

	config.Postgres.Table("assets").Select("assets.id, assets.inscription_id, assets.address, assets.type, assets.category, assets.collection, assets.tag").Joins("left join inscriptions on assets.id=inscriptions.id").Where(w).Order("id").Find(&assets)

	var brc420AssetsWithName []model.BRC420EntryWithName
	config.Postgres.Table("brc420_entries").
		Select("brc420_entries.id", "brc420_entries.inscription_id", "brc420_entries.address", "brc420_entries.content_type", "brc420_entries.ref", "brc420_entries.fee_at", "brc420_entries.mint_timestamp", "brc420_details.name").
		Joins("left join brc420_details on brc420_entries.ref=brc420_details.tick").Where("brc420_entries.address = ? and content_type in ?", s.Address, []string{"text/html", "text/html;charset=utf-8"}).Find(&brc420AssetsWithName)

	findAllAssets := []string{}
	for _, asset := range brc420AssetsWithName {
		findAllAssets = append(findAllAssets, asset.Ref)
	}
	var brc420Assets []model.Asset
	ow := "assets.id >= 0 and inscriptions.oss_url is not null and assets.inscription_id in ?"

	if s.Type != "" {
		ow = fmt.Sprintf("%s and type in ('%s', 'character')", ow, s.Type)
	}

	if s.Category != "" {
		ow = fmt.Sprintf("%s and category = '%s'", ow, s.Category)
	}

	if s.Collection != "" {
		ow = fmt.Sprintf("%s and collection = '%s'", ow, s.Collection)
	}

	if s.Tag != "" {
		ow = fmt.Sprintf("%s and tag like '%%%s%%'", ow, s.Tag)
	}

	config.Postgres.Debug().Table("assets").Select("assets.id, assets.inscription_id, assets.address, assets.type, assets.category, assets.collection, assets.tag").Joins("left join inscriptions on assets.id=inscriptions.id").Where(ow, findAllAssets).Order("id").Find(&brc420Assets)

	assets = append(assets, brc420Assets...)

	sort.Slice(assets, func(i, j int) bool {
		return assets[i].Id < assets[j].Id
	})
	cnt := len(assets)

	last := s.Offset + s.Limit
	if last > len(assets) {
		last = len(assets)
	}

	if s.Offset < len(assets) {
		assets = assets[s.Offset:last]
	} else {
		assets = make([]model.Asset, 0)
	}

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildAssetsListWithCntResponse(int64(cnt), assets),
	}
}

func (s *AssetsListService) ListWithOssAndBRC420V2() serializer.Response {
	var assets []model.Asset

	w := fmt.Sprintf("assets.address = '%s' and assets.id >= 0 and inscriptions.oss_url is not null", s.Address)
	if s.Type != "" {
		w = fmt.Sprintf("%s and type in ('%s', 'character')", w, s.Type)
	}

	if s.Category != "" {
		w = fmt.Sprintf("%s and category = '%s'", w, s.Category)
	}

	if s.Collection != "" {
		w = fmt.Sprintf("%s and collection = '%s'", w, s.Collection)
	}

	if s.Tag != "" {
		w = fmt.Sprintf("%s and tag like '%%%s%%'", w, s.Tag)
	}

	config.Postgres.Table("assets").Select("assets.id, assets.inscription_id, assets.address, assets.type, assets.category, assets.collection, assets.tag").Joins("left join inscriptions on assets.id=inscriptions.id").Where(w).Order("id").Find(&assets)

	var brc420AssetsWithName []model.BRC420EntryWithName
	config.Postgres.Table("brc420_entries").
		Select("brc420_entries.id", "brc420_entries.inscription_id", "brc420_entries.address", "brc420_entries.content_type", "brc420_entries.ref", "brc420_entries.fee_at", "brc420_entries.mint_timestamp", "brc420_details.name").
		Joins("left join brc420_details on brc420_entries.ref=brc420_details.tick").Where("brc420_entries.address = ? and content_type in ?", s.Address, []string{"text/html", "text/html;charset=utf-8"}).Find(&brc420AssetsWithName)

	findAllAssets := []string{}
	for _, asset := range brc420AssetsWithName {
		findAllAssets = append(findAllAssets, asset.Ref)
	}
	var brc420Assets []model.Asset
	ow := "assets.id >= 0 and inscriptions.oss_url is not null and assets.inscription_id in ?"

	if s.Type != "" {
		ow = fmt.Sprintf("%s and type in ('%s', 'character')", ow, s.Type)
	}

	if s.Category != "" {
		ow = fmt.Sprintf("%s and category = '%s'", ow, s.Category)
	}

	if s.Collection != "" {
		ow = fmt.Sprintf("%s and collection = '%s'", ow, s.Collection)
	}

	if s.Tag != "" {
		ow = fmt.Sprintf("%s and tag like '%%%s%%'", ow, s.Tag)
	}

	config.Postgres.Table("assets").Select("assets.id, assets.inscription_id, assets.address, assets.type, assets.category, assets.collection, assets.tag").Joins("left join inscriptions on assets.id=inscriptions.id").Where(ow, findAllAssets).Order("id").Find(&brc420Assets)

	assets = append(assets, brc420Assets...)

	sort.Slice(assets, func(i, j int) bool {
		return assets[i].Id < assets[j].Id
	})
	cnt := len(assets)

	dlcService := AssetsListService{
		Address:    s.Address,
		Type:       "",
		Category:   "",
		Collection: "",
		Tag:        "",
		Offset:     0,
		Limit:      10000,
	}
	var as []model.AssDLC = []model.AssDLC{}
	for _, a := range assets {
		as = append(as, model.AssDLC{
			Asset: a,
			DLC:   false,
		})
	}
	dlcRes := dlcService.ListDLCWithOssAndBRC420()
	dlcList, ok := dlcRes.Data.(*serializer.AssetsListResponse)
	if ok {
		for _, item := range dlcList.Items {
			if cache.DLCToAssets[item.InscriptionId] != nil {
				// assets = append(assets, cache.DLCToAssets[item.InscriptionId]...)
				// cnt += len(cache.DLCToAssets[item.InscriptionId])
				for _, ass := range cache.DLCToAssets[item.InscriptionId] {
					if ass.Category == s.Category && s.Category != "" {
						as = append(as, model.AssDLC{
							Asset: ass,
							DLC:   true,
						})
						cnt += 1
					} else if ass.Collection == s.Collection && s.Collection != "" {
						as = append(as, model.AssDLC{
							Asset: ass,
							DLC:   true,
						})
						cnt += 1
					} else if ass.Tag == s.Tag && s.Tag != "" {
						as = append(as, model.AssDLC{
							Asset: ass,
							DLC:   true,
						})
						cnt += 1
					} else if ass.Type == s.Type && s.Type != "" {
						as = append(as, model.AssDLC{
							Asset: ass,
							DLC:   true,
						})
						cnt += 1
					}
				}
			}
		}
	}

	sort.Slice(as, func(i, j int) bool {
		return as[i].Id < as[j].Id
	})
	cnt = len(as)

	last := s.Offset + s.Limit
	if last > len(as) {
		last = len(as)
	}

	if s.Offset < len(as) {
		as = as[s.Offset:last]
	} else {
		as = make([]model.AssDLC, 0)
	}

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildAssetsListWithCntResponseDLC(int64(cnt), as),
	}
}

func (s *AssetsListService) ListWithOssAndBRC420V3() serializer.Response {
	var assets []model.Asset

	w := fmt.Sprintf("assets.address = '%s' and assets.id >= 0 and inscriptions.oss_url is not null", s.Address)
	if s.Type != "" {
		w = fmt.Sprintf("%s and type in ('%s', 'character')", w, s.Type)
	}

	if s.Category != "" {
		w = fmt.Sprintf("%s and category = '%s'", w, s.Category)
	}

	if s.Collection != "" {
		w = fmt.Sprintf("%s and collection = '%s'", w, s.Collection)
	}

	if s.Tag != "" {
		w = fmt.Sprintf("%s and tag like '%%%s%%'", w, s.Tag)
	}

	config.Postgres.Table("assets").Select("assets.id, assets.inscription_id, assets.address, assets.type, assets.category, assets.collection, assets.tag").Joins("left join inscriptions on assets.id=inscriptions.id").Where(w).Order("id").Find(&assets)

	var brc420AssetsWithName []model.BRC420EntryWithName
	config.Postgres.Table("brc420_entries").
		Select("brc420_entries.id", "brc420_entries.inscription_id", "brc420_entries.address", "brc420_entries.content_type", "brc420_entries.ref", "brc420_entries.fee_at", "brc420_entries.mint_timestamp", "brc420_details.name").
		Joins("left join brc420_details on brc420_entries.ref=brc420_details.tick").Where("brc420_entries.address = ? and content_type in ?", s.Address, []string{"text/html", "text/html;charset=utf-8"}).Find(&brc420AssetsWithName)

	findAllAssets := []string{}
	for _, asset := range brc420AssetsWithName {
		findAllAssets = append(findAllAssets, asset.Ref)
	}
	var brc420Assets []model.Asset
	ow := "assets.id >= 0 and inscriptions.oss_url is not null and assets.inscription_id in ?"

	if s.Type != "" {
		ow = fmt.Sprintf("%s and type in ('%s', 'character')", ow, s.Type)
	}

	if s.Category != "" {
		ow = fmt.Sprintf("%s and category = '%s'", ow, s.Category)
	}

	if s.Collection != "" {
		ow = fmt.Sprintf("%s and collection = '%s'", ow, s.Collection)
	}

	if s.Tag != "" {
		ow = fmt.Sprintf("%s and tag like '%%%s%%'", ow, s.Tag)
	}

	config.Postgres.Table("assets").Select("assets.id, assets.inscription_id, assets.address, assets.type, assets.category, assets.collection, assets.tag").Joins("left join inscriptions on assets.id=inscriptions.id").Where(ow, findAllAssets).Order("id").Find(&brc420Assets)

	assets = append(assets, brc420Assets...)

	sort.Slice(assets, func(i, j int) bool {
		return assets[i].Id < assets[j].Id
	})
	cnt := len(assets)

	dlcService := AssetsListService{
		Address:    s.Address,
		Type:       "",
		Category:   "",
		Collection: "",
		Tag:        "",
		Offset:     0,
		Limit:      10000,
	}
	var as []model.AssDLC = []model.AssDLC{}
	for _, a := range assets {
		as = append(as, model.AssDLC{
			Asset: a,
			DLC:   false,
		})
	}
	dlcRes := dlcService.ListDLCWithOssAndBRC420V3()
	dlcList, ok := dlcRes.Data.(*serializer.AssetsListResponse)
	if ok {
		for _, item := range dlcList.Items {
			if item.InscriptionId == "e80eadea13e2175949168c279f7a47a467b6c5e00a84d45dd8ae40aefda89fe5i0" {
				item.InscriptionId = "dbdbc1ff9fa94d1149c240f742fa444a853b0101d3fb898376e00adb8792454di0"
			}
			if cache.DLCToAssets[item.InscriptionId] != nil {
				// assets = append(assets, cache.DLCToAssets[item.InscriptionId]...)
				// cnt += len(cache.DLCToAssets[item.InscriptionId])
				for _, ass := range cache.DLCToAssets[item.InscriptionId] {
					if ass.Category == s.Category && s.Category != "" {
						as = append(as, model.AssDLC{
							Asset: ass,
							DLC:   true,
						})
						cnt += 1
					} else if ass.Collection == s.Collection && s.Collection != "" {
						as = append(as, model.AssDLC{
							Asset: ass,
							DLC:   true,
						})
						cnt += 1
					} else if ass.Tag == s.Tag && s.Tag != "" {
						as = append(as, model.AssDLC{
							Asset: ass,
							DLC:   true,
						})
						cnt += 1
					} else if ass.Type == s.Type && s.Type != "" {
						as = append(as, model.AssDLC{
							Asset: ass,
							DLC:   true,
						})
						cnt += 1
					}
				}
			}
		}
	}

	sort.Slice(as, func(i, j int) bool {
		return as[i].Id < as[j].Id
	})
	cnt = len(as)

	last := s.Offset + s.Limit
	if last > len(as) {
		last = len(as)
	}

	if s.Offset < len(as) {
		as = as[s.Offset:last]
	} else {
		as = make([]model.AssDLC, 0)
	}

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildAssetsListWithCntResponseDLC(int64(cnt), as),
	}
}

func (s *AssetsListService) ListDLCWithOssAndBRC420() serializer.Response {
	var assets []model.Asset

	w := fmt.Sprintf("assets.address = '%s' and assets.id >= 0 and inscriptions.oss_url is not null", s.Address)
	w = fmt.Sprintf("%s and type in ('%s', 'character')", w, s.Type)

	w = fmt.Sprintf("%s and category = '%s'", w, s.Category)

	w = fmt.Sprintf("%s and collection = '%s'", w, s.Collection)

	w = fmt.Sprintf("%s and tag like '%%%s%%'", w, s.Tag)

	config.Postgres.Table("assets").Select("assets.id, assets.inscription_id, assets.address, assets.type, assets.category, assets.collection, assets.tag").Joins("left join inscriptions on assets.id=inscriptions.id").Where(w).Order("id").Find(&assets)

	var brc420AssetsWithName []model.BRC420EntryWithName
	config.Postgres.Table("brc420_entries").
		Select("brc420_entries.id", "brc420_entries.inscription_id", "brc420_entries.address", "brc420_entries.content_type", "brc420_entries.ref", "brc420_entries.fee_at", "brc420_entries.mint_timestamp", "brc420_details.name").
		Joins("left join brc420_details on brc420_entries.ref=brc420_details.tick").Where("brc420_entries.address = ? and content_type in ?", s.Address, []string{"text/html", "text/html;charset=utf-8"}).Find(&brc420AssetsWithName)

	findAllAssets := []string{}
	for _, asset := range brc420AssetsWithName {
		findAllAssets = append(findAllAssets, asset.Ref)
	}
	var brc420Assets []model.Asset
	ow := "assets.id >= 0 and inscriptions.oss_url is not null and assets.inscription_id in ?"

	ow = fmt.Sprintf("%s and type in ('%s', 'character')", ow, s.Type)

	ow = fmt.Sprintf("%s and category = '%s'", ow, s.Category)

	ow = fmt.Sprintf("%s and collection = '%s'", ow, s.Collection)

	ow = fmt.Sprintf("%s and tag like '%%%s%%'", ow, s.Tag)

	config.Postgres.Table("assets").Select("assets.id, assets.inscription_id, assets.address, assets.type, assets.category, assets.collection, assets.tag").Joins("left join inscriptions on assets.id=inscriptions.id").Where(ow, findAllAssets).Order("id").Find(&brc420Assets)

	assets = append(assets, brc420Assets...)

	sort.Slice(assets, func(i, j int) bool {
		return assets[i].Id < assets[j].Id
	})
	cnt := len(assets)

	last := s.Offset + s.Limit
	if last > len(assets) {
		last = len(assets)
	}

	if s.Offset < len(assets) {
		assets = assets[s.Offset:last]
	} else {
		assets = make([]model.Asset, 0)
	}

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildAssetsListWithCntResponse(int64(cnt), assets),
	}
}

func (s *AssetsListService) ListDLCWithOssAndBRC420V3() serializer.Response {
	var assets []model.Asset

	w := fmt.Sprintf("assets.address = '%s' and assets.id >= 0 and inscriptions.oss_url is not null", s.Address)
	w = fmt.Sprintf("%s and type in ('%s', 'character')", w, s.Type)

	w = fmt.Sprintf("%s and category = '%s'", w, s.Category)

	w = fmt.Sprintf("%s and collection = '%s'", w, s.Collection)

	w = fmt.Sprintf("%s and tag like '%%%s%%'", w, s.Tag)

	config.Postgres.Table("assets").Select("assets.id, assets.inscription_id, assets.address, assets.type, assets.category, assets.collection, assets.tag").Joins("left join inscriptions on assets.id=inscriptions.id").Where(w).Order("id").Find(&assets)

	var brc420AssetsWithName []model.BRC420EntryWithName
	config.Postgres.Debug().Table("brc420_entries").
		Select("brc420_entries.id", "brc420_entries.inscription_id", "brc420_entries.address", "brc420_entries.content_type", "brc420_entries.ref", "brc420_entries.fee_at", "brc420_entries.mint_timestamp", "brc420_details.name").
		Joins("left join brc420_details on brc420_entries.ref=brc420_details.tick").Where("(brc420_entries.address = ? and content_type in ?) or brc420_entries.ref = 'e80eadea13e2175949168c279f7a47a467b6c5e00a84d45dd8ae40aefda89fe5i0'", s.Address, []string{"text/html", "text/html;charset=utf-8"}).Find(&brc420AssetsWithName)

	findAllAssets := []string{}
	for _, asset := range brc420AssetsWithName {
		if asset.Ref == "e80eadea13e2175949168c279f7a47a467b6c5e00a84d45dd8ae40aefda89fe5i0" {
			asset.Ref = "dbdbc1ff9fa94d1149c240f742fa444a853b0101d3fb898376e00adb8792454di0"
		}
		findAllAssets = append(findAllAssets, asset.Ref)
	}
	var brc420Assets []model.Asset
	ow := "assets.id >= 0 and inscriptions.oss_url is not null and assets.inscription_id in ?"

	ow = fmt.Sprintf("%s and type in ('%s', 'character')", ow, s.Type)

	ow = fmt.Sprintf("%s and category = '%s'", ow, s.Category)

	ow = fmt.Sprintf("%s and collection = '%s'", ow, s.Collection)

	ow = fmt.Sprintf("%s and tag like '%%%s%%'", ow, s.Tag)

	config.Postgres.Debug().Table("assets").Select("assets.id, assets.inscription_id, assets.address, assets.type, assets.category, assets.collection, assets.tag").Joins("left join inscriptions on assets.id=inscriptions.id").Where(ow, findAllAssets).Order("id").Find(&brc420Assets)

	assets = append(assets, brc420Assets...)

	sort.Slice(assets, func(i, j int) bool {
		return assets[i].Id < assets[j].Id
	})
	cnt := len(assets)

	last := s.Offset + s.Limit
	if last > len(assets) {
		last = len(assets)
	}

	if s.Offset < len(assets) {
		assets = assets[s.Offset:last]
	} else {
		assets = make([]model.Asset, 0)
	}

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildAssetsListWithCntResponse(int64(cnt), assets),
	}
}
