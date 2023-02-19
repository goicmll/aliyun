package aliyun

import (
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

// ###### ECS
type EcsInstanceClient struct {
	AliYunClient aliYunClient
}

func (receiver *EcsInstanceClient) SetEndPoint(endPoint ...string) {
	if endPoint == nil {
		receiver.AliYunClient.Client.Endpoint = tea.String("ecs-cn-hangzhou.aliyuncs.com")
	} else {
		receiver.AliYunClient.Client.Endpoint = &endPoint[0]
	}

}

func (receiver EcsInstanceClient) GetEcsInstanceInfoList(regionID string, pageNum, pageSize int32) (*ecs20140526.DescribeInstancesResponseBody, error) {
	describeInstancesRequest := &ecs20140526.DescribeInstancesRequest{
		RegionId:   tea.String(regionID),
		PageNumber: tea.Int32(pageNum),
		PageSize:   tea.Int32(pageSize),
	}
	client := ecs20140526.Client{Client: *receiver.AliYunClient.Client}
	result, err := client.DescribeInstances(describeInstancesRequest)
	if err != nil {
		return nil, AliCloudError{err.Error()}
	}
	return result.Body, nil
}
