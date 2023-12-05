package serializer

import (
	"server-template/cache"
	"server-template/model"
)

type AssetCollectionResponse struct {
	Name           string   `json:"name"`
	CollectionName string   `json:"collection_name"`
	InscriptionIds []string `json:"inscription_ids"`
	Type           string   `json:"type"`
}

type AssetCollectionListResponse struct {
	Count int64                      `json:"count"`
	Items []*AssetCollectionResponse `json:"items"`
}

func BuildAssetCollectionListResponse(brc420Details []model.BRC420Detail) AssetCollectionListResponse {
	var list AssetCollectionListResponse

	list.Count = int64(len(brc420Details))
	for _, detail := range brc420Details {
		res := &AssetCollectionResponse{
			Name: detail.Name,
		}

		if detail.Tick == "e80eadea13e2175949168c279f7a47a467b6c5e00a84d45dd8ae40aefda89fe5i0" {
			res.CollectionName = cache.DLCToAssets["dbdbc1ff9fa94d1149c240f742fa444a853b0101d3fb898376e00adb8792454di0"][0].Collection
		} else if detail.Tick == "6fdd49dc8e9de8faa3ecb8a81ef5a9b81f77c2d76f14ec5a94d62229e586be00i0" {
			res.CollectionName = "BitmapKingdomDLCBitmapTown"
		} else {
			if len(cache.DLCToAssets[detail.Tick]) != 0 {
				res.CollectionName = cache.DLCToAssets[detail.Tick][0].Collection
			}
		}
		list.Items = append(list.Items, res)
	}

	list.Items = append(list.Items, &AssetCollectionResponse{
		Name:           "others",
		CollectionName: "others",
	})

	return list
}

func BuildAssetCollectionNameListResponse(brc420Details []model.BRC420Detail, nameToInsids map[string][]string) AssetCollectionListResponse {
	var list AssetCollectionListResponse

	list.Count = int64(len(brc420Details) + len(nameToInsids))
	for _, detail := range brc420Details {
		res := &AssetCollectionResponse{
			Name:           detail.Name,
			InscriptionIds: []string{detail.Tick},
			Type:           "dlc",
		}

		if detail.Tick == "e80eadea13e2175949168c279f7a47a467b6c5e00a84d45dd8ae40aefda89fe5i0" {
			res.CollectionName = cache.DLCToAssets["dbdbc1ff9fa94d1149c240f742fa444a853b0101d3fb898376e00adb8792454di0"][0].Collection
		} else if detail.Tick == "6fdd49dc8e9de8faa3ecb8a81ef5a9b81f77c2d76f14ec5a94d62229e586be00i0" {
			res.CollectionName = "BitmapKingdomDLCBitmapTown"
		} else {
			if len(cache.DLCToAssets[detail.Tick]) != 0 {
				res.CollectionName = cache.DLCToAssets[detail.Tick][0].Collection
			}
		}
		list.Items = append(list.Items, res)
	}

	for k, v := range nameToInsids {
		res := &AssetCollectionResponse{
			Name:           k,
			InscriptionIds: v,
			Type:           "brc420",
		}

		list.Items = append(list.Items, res)
	}

	list.Items = append(list.Items, &AssetCollectionResponse{
		Name:           "others",
		CollectionName: "others",
		InscriptionIds: []string{},
		Type:           "others",
	})

	return list
}
