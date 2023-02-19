package aliyun

import (
	rds20140815 "github.com/alibabacloud-go/rds-20140815/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

// #############RDS
type RDSInstanceClient struct {
	AliYunClient aliYunClient
}

// 设置端点
func (receiver *RDSInstanceClient) SetEndPoint(endPoint ...string) {
	if endPoint == nil {
		receiver.AliYunClient.Client.Endpoint = tea.String("rds.aliyuncs.com")
	} else {
		receiver.AliYunClient.Client.Endpoint = &endPoint[0]
	}

}

// 获取实例列表
func (receiver RDSInstanceClient) GetRDSInstanceInfoList(regionID string, pageNum, pageSize int32) (*rds20140815.DescribeDBInstancesResponseBody, error) {
	describeDBInstancesRequest := &rds20140815.DescribeDBInstancesRequest{
		RegionId:   tea.String(regionID),
		PageNumber: tea.Int32(pageNum),
		PageSize:   tea.Int32(pageSize),
	}
	client := rds20140815.Client{Client: *receiver.AliYunClient.Client}
	result, err := client.DescribeDBInstances(describeDBInstancesRequest)
	if err != nil {
		return nil, AliCloudError{err.Error()}
	}
	return result.Body, nil
}

func (receiver RDSInstanceClient) GetDatabaseList(instanceID string, pageNum, pageSize int32) (*rds20140815.DescribeDatabasesResponseBody, error) {
	describeDatabasesRequest := &rds20140815.DescribeDatabasesRequest{
		DBInstanceId: tea.String(instanceID),
		PageNumber:   tea.Int32(pageNum),
		PageSize:     tea.Int32(pageSize),
	}
	client := rds20140815.Client{Client: *receiver.AliYunClient.Client}
	result, err := client.DescribeDatabases(describeDatabasesRequest)
	if err != nil {
		return nil, AliCloudError{err.Error()}
	}
	return result.Body, nil
}
