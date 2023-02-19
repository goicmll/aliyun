package aliyun

import (
	dds20151201 "github.com/alibabacloud-go/dds-20151201/client"
	"github.com/alibabacloud-go/tea/tea"
)

// #############mongodb
type MongoClient struct {
	AliYunClient aliYunClient
}

// 设置端点
func (receiver *MongoClient) SetEndPoint(endPoint ...string) {
	if endPoint == nil {
		receiver.AliYunClient.Client.Endpoint = tea.String("mongodb.aliyuncs.com")
	} else {
		receiver.AliYunClient.Client.Endpoint = &endPoint[0]
	}

}

// pageSize取值： 30、50、100，默认值为30
func (receiver MongoClient) GetMongoInstanceInfoList(regionID string, pageNum, pageSize int32) (*dds20151201.DescribeDBInstancesResponseBody, error) {
	queryMongoListRequest := &dds20151201.DescribeDBInstancesRequest{
		RegionId:   tea.String(regionID),
		PageNumber: tea.Int32(pageNum),
		PageSize:   tea.Int32(pageSize),
	}
	client := dds20151201.Client{Client: *receiver.AliYunClient.Client}
	result, err := client.DescribeDBInstances(queryMongoListRequest)
	if err != nil {
		return nil, AliCloudError{err.Error()}
	}
	return result.Body, nil
}

func (receiver MongoClient) GetMongoInstanceDetail(instanceID string) (*dds20151201.DescribeDBInstanceAttributeResponseBody, error) {
	mongoAttributeRequest := &dds20151201.DescribeDBInstanceAttributeRequest{
		DBInstanceId: tea.String(instanceID),
	}
	client := dds20151201.Client{Client: *receiver.AliYunClient.Client}
	result, err := client.DescribeDBInstanceAttribute(mongoAttributeRequest)
	if err != nil {
		return nil, AliCloudError{err.Error()}
	}
	return result.Body, nil
}
