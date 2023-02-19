package aliyun

import (
	"fmt"

	ons20190214 "github.com/alibabacloud-go/ons-20190214/client"
	"github.com/alibabacloud-go/tea/tea"
)

// #############Redis
type RocketMQInstanceClient struct {
	AliYunClient aliYunClient
}

// 设置端点
func (receiver *RocketMQInstanceClient) SetEndPoint(endPoint ...string) {
	if endPoint == nil {
		receiver.AliYunClient.Client.Endpoint = tea.String("ons.cn-hangzhou.aliyuncs.com")
	} else {
		receiver.AliYunClient.Client.Endpoint = &endPoint[0]
	}
}

// 设置客户端协议, mq这边有些https不能用, 要设置成http
func (receiver *RocketMQInstanceClient) SetProtocol(protocol ...string) {
	if protocol == nil {
		receiver.AliYunClient.Client.Protocol = tea.String("https")
	} else {
		receiver.AliYunClient.Client.Protocol = &protocol[0]
	}
}

// 通过区域id 合成端点
func (receiver *RocketMQInstanceClient) generateEndPointByRegionID(regionID string) string {
	return fmt.Sprintf("ons.%s.aliyuncs.com", regionID)
}

// 获取列表
func (receiver RocketMQInstanceClient) GetRocketMQInstanceInfoList(regionID string) (*ons20190214.OnsInstanceInServiceListResponseBody, error) {
	onsInstanceInServiceListRequest := &ons20190214.OnsInstanceInServiceListRequest{}
	receiver.SetEndPoint(receiver.generateEndPointByRegionID(regionID))
	client := ons20190214.Client{Client: *receiver.AliYunClient.Client}
	result, err := client.OnsInstanceInServiceList(onsInstanceInServiceListRequest)
	if err != nil {
		return nil, AliCloudError{err.Error()}
	}
	return result.Body, nil
}

// 获取详情
func (receiver RocketMQInstanceClient) GetRocketMQInstanceDetail(regionID, instanceID string) (*ons20190214.OnsInstanceBaseInfoResponseBody, error) {
	onsInstanceBaseInfoRequest := &ons20190214.OnsInstanceBaseInfoRequest{
		InstanceId: tea.String(instanceID),
	}
	receiver.SetEndPoint(receiver.generateEndPointByRegionID(regionID))
	client := ons20190214.Client{Client: *receiver.AliYunClient.Client}
	result, err := client.OnsInstanceBaseInfo(onsInstanceBaseInfoRequest)
	if err != nil {
		return nil, AliCloudError{err.Error()}
	}
	return result.Body, nil
}
