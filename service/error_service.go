package service

import (
	"fmt"

	"server-template/serializer"
)

type ErrorService struct {
	Log string `form:"log" json:"log" binding:"required"`
	Err string `form:"err" json:"err" binding:"required"`
}

func (service *ErrorService) Println() serializer.Response {
	fmt.Println(fmt.Sprintf("%s: %s", service.Log, service.Err))
	return serializer.Response{
		Code: 200,
	}
}
