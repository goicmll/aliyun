package aliyun

import (
	domain20180129 "github.com/alibabacloud-go/domain-20180129/client"
	"github.com/alibabacloud-go/tea/tea"
)

// #############域名获取
type DomainClient struct {
	AliYunClient aliYunClient
}

// 设置端点
func (receiver *DomainClient) SetEndPoint(endPoint ...string) {
	if endPoint == nil {
		receiver.AliYunClient.Client.Endpoint = tea.String("domain.aliyuncs.com")
	} else {
		receiver.AliYunClient.Client.Endpoint = &endPoint[0]
	}

}

func (receiver DomainClient) GetRocketMQInstanceInfoList(pageNum, pageSize int32) (*domain20180129.QueryDomainListResponseBody, error) {
	queryDomainListRequest := &domain20180129.QueryDomainListRequest{
		PageNum:  tea.Int32(pageNum),
		PageSize: tea.Int32(pageSize),
	}
	client := domain20180129.Client{Client: *receiver.AliYunClient.Client}
	result, err := client.QueryDomainList(queryDomainListRequest)
	if err != nil {
		return nil, AliCloudError{err.Error()}
	}
	return result.Body, nil
}
