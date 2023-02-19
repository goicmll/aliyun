package aliyun

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

//### 基本用法
// 通过NewAliYunClient 创建aliyun clients
// 再通过各个资源的 struct 后去各个资源特有的client
// oss 的除外

// ###### 客户端
type aliYunClient struct {
	accessKeyID     string
	accessKeySecret string
	Client          *openapi.Client
}

// 客户端登录
func (receiver *aliYunClient) login() {
	config := &openapi.Config{
		AccessKeyId:     tea.String(receiver.accessKeyID),
		AccessKeySecret: tea.String(receiver.accessKeySecret),
		ReadTimeout:     tea.Int(7000),
		ConnectTimeout:  tea.Int(3000),
	}
	result, _ := ecs20140526.NewClient(config)
	receiver.Client = &result.Client
}

// 创建已登录的客户端
func NewAliYunClient(ak, akSecret string) aliYunClient {
	client := &aliYunClient{accessKeyID: ak, accessKeySecret: akSecret}
	client.login()
	return *client
}
