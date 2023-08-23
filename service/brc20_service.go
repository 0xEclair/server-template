package service

import (
	"server-template/serializer"
	"server-template/third/bestinslot"
)

type BRC20Service struct {
	Address string `json:"address" form:"address" binding:"required"`
}

func (s *BRC20Service) AllBalance() serializer.Response {
	res, err := bestinslot.BRC20ByAddress(s.Address)
	if err != nil {
		return serializer.Response{
			Code: 200,
			Data: nil,
			Err:  err.Error(),
		}
	}

	return serializer.Response{
		Code: 200,
		Data: res,
	}
}
