package aliyun

import (
	r_kvstore20150101 "github.com/alibabacloud-go/r-kvstore-20150101/client"
	"github.com/alibabacloud-go/tea/tea"
)

// #############Redis
type RedisInstanceClient struct {
	AliYunClient aliYunClient
}

// 设置端点
func (receiver *RedisInstanceClient) SetEndPoint(endPoint ...string) {
	if endPoint == nil {
		receiver.AliYunClient.Client.Endpoint = tea.String("r-kvstore.aliyuncs.com")
	} else {
		receiver.AliYunClient.Client.Endpoint = &endPoint[0]
	}
}

// 获取列表
func (receiver RedisInstanceClient) GetRedisInstanceInfoList(regionID string, pageNum, pageSize int32) (*r_kvstore20150101.DescribeInstancesResponseBody, error) {
	describeInstancesRequest := &r_kvstore20150101.DescribeInstancesRequest{
		PageNumber: tea.Int32(pageNum),
		PageSize:   tea.Int32(pageSize),
		RegionId:   tea.String(regionID),
	}
	client := r_kvstore20150101.Client{Client: *receiver.AliYunClient.Client}
	result, err := client.DescribeInstances(describeInstancesRequest)
	if err != nil {
		return nil, AliCloudError{err.Error()}
	}
	return result.Body, nil
}
