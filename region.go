package aliyun

import (
	dds20151201 "github.com/alibabacloud-go/dds-20151201/client"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v2/client"
	ons20190214 "github.com/alibabacloud-go/ons-20190214/client"
	r_kvstore20150101 "github.com/alibabacloud-go/r-kvstore-20150101/client"
	rds20140815 "github.com/alibabacloud-go/rds-20140815/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

// ######### 地域
type RegionClient struct {
	AliYunClient aliYunClient
}

// 获取ecs区域列表
func (receiver *RegionClient) GetEcsRegionInfoList() (*ecs20140526.DescribeRegionsResponseBody, error) {
	describeRegionsRequest := &ecs20140526.DescribeRegionsRequest{}
	receiver.AliYunClient.Client.Endpoint = tea.String("ecs-cn-hangzhou.aliyuncs.com")
	client := ecs20140526.Client{Client: *receiver.AliYunClient.Client}
	result, err := client.DescribeRegions(describeRegionsRequest)
	if err != nil {
		return nil, AliCloudError{err.Error()}
	}
	return result.Body, nil
}

// 获取rds区域列表
func (receiver *RegionClient) GetRdsRegionInfoList() (*rds20140815.DescribeRegionsResponseBody, error) {
	describeRegionsRequest := &rds20140815.DescribeRegionsRequest{}
	receiver.AliYunClient.Client.Endpoint = tea.String("rds.aliyuncs.com")
	client := rds20140815.Client{Client: *receiver.AliYunClient.Client}
	result, err := client.DescribeRegions(describeRegionsRequest)
	if err != nil {
		return nil, AliCloudError{err.Error()}
	}
	return result.Body, nil
}

// 获取redis区域列表
func (receiver *RegionClient) GetRedisRegionInfoList() (*r_kvstore20150101.DescribeRegionsResponseBody, error) {
	describeRegionsRequest := &r_kvstore20150101.DescribeRegionsRequest{}
	receiver.AliYunClient.Client.Endpoint = tea.String("r-kvstore.aliyuncs.com")
	client := r_kvstore20150101.Client{Client: *receiver.AliYunClient.Client}
	result, err := client.DescribeRegions(describeRegionsRequest)
	if err != nil {
		return nil, AliCloudError{err.Error()}
	}
	return result.Body, nil
}

// 获取rocket mq区域列表
func (receiver *RegionClient) GetRocketMQRegionInfoList() (*ons20190214.OnsRegionListResponseBody, error) {
	receiver.AliYunClient.Client.Endpoint = tea.String("ons.cn-hangzhou.aliyuncs.com")
	client := ons20190214.Client{Client: *receiver.AliYunClient.Client}
	result, err := client.OnsRegionList()
	if err != nil {
		return nil, AliCloudError{err.Error()}
	}
	return result.Body, nil
}

// 获取 mongo 区域列表
func (receiver *RegionClient) GetMongoRegionInfoList() (*dds20151201.DescribeRegionsResponseBody, error) {
	describeRegionsRequest := &dds20151201.DescribeRegionsRequest{}
	receiver.AliYunClient.Client.Endpoint = tea.String("mongodb.aliyuncs.com")
	client := dds20151201.Client{Client: *receiver.AliYunClient.Client}
	result, err := client.DescribeRegions(describeRegionsRequest)
	if err != nil {
		return nil, AliCloudError{err.Error()}
	}
	return result.Body, nil
}
