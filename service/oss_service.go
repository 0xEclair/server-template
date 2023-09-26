package service

import (
	"os"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"

	"server-template/serializer"
)

type OssService struct {
}

type CreateKeyResponse struct {
	AccessKey    string `json:"access_key"`
	AccessSecret string `json:"access_secret"`
}

func (s *OssService) CreateKey() serializer.Response {
	client, err := sts.NewClientWithAccessKey(os.Getenv("oss_access_region"), os.Getenv("oss_access_key"), os.Getenv("oss_access_secret"))
	if err != nil {
		return serializer.Response{
			Code: 201,
			Err:  "Bad Request.",
		}
	}

	request := sts.CreateAssumeRoleRequest()
	request.Scheme = "https"

	request.RoleArn = os.Getenv("oss_access_arn")
	request.RoleSessionName = os.Getenv("oss_access_name")
	response, err := client.AssumeRole(request)
	if err != nil {
		return serializer.Response{
			Code: 202,
			Err:  "Bad Request.",
		}
	}

	return serializer.Response{
		Code: 200,
		Data: CreateKeyResponse{
			AccessKey:    response.Credentials.AccessKeyId,
			AccessSecret: response.Credentials.AccessKeySecret,
		},
	}
}
