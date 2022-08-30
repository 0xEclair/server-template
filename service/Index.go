package service

import (
	"fmt"

	"server-template/serializer"
)

type IndexService struct {
	Log string `form:"log" json:"log" binding:"required"`
	Err string `form:"err" json:"err" binding:"required"`
}

func (service *IndexService) Accept() serializer.Response {
	fmt.Println(fmt.Sprintf("%s: %s", service.Log, service.Err))
	return serializer.Response{
		Code: 200,
	}
}
