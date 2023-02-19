package aliyun

import (
	"fmt"
	"regexp"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type OSSClient struct {
	Endpoint        string
	AccessKeyID     string
	AccessKeySecret string
	UseCname        bool
	client          *oss.Client
}

// 新建一个客端
func NewOSSClient(endpoint, accessKeyID, accessKeySecret string, useCname bool) (*OSSClient, error) {
	client, err := oss.New(endpoint, accessKeyID, accessKeySecret, oss.UseCname(useCname))
	oc := OSSClient{
		Endpoint:        endpoint,
		AccessKeyID:     accessKeyID,
		AccessKeySecret: accessKeySecret,
		UseCname:        useCname,
		client:          client,
	}
	if err != nil {
		return nil, AliCloudError{err.Error()}
	}
	return &oc, nil
}

// 判断bucket是否存在
func (receiver OSSClient) IsBucketExist(bucketName string) (bool, error) {
	isExist, err := receiver.client.IsBucketExist(bucketName)
	if err != nil {
		return false, AliCloudError{err.Error()}
	}
	return isExist, nil

}

// 上传文件
// remoteFilePath 为在oss上相对于bucket的路径
func (receiver OSSClient) UploadFile(bucketName, remoteFilePath, localFilePath string, partSize int64) error {
	isExist, err := receiver.IsBucketExist(bucketName)
	if err != nil {
		return AliCloudError{err.Error()}
	}
	if !isExist {
		return AliCloudError{fmt.Sprintf("bucket(%s)不存在!", bucketName)}
	}
	filePathRegex := "^[A-Za-z0-9]+(/{1}[A-Za-z0-9_-]+)*/[A-Za-z0-9.@_-]+$"

	reg, _ := regexp.Compile(filePathRegex)
	if !reg.Match([]byte(remoteFilePath)) {
		return AliCloudError{fmt.Sprintf("文件路径(%s)不符合规范(%s)!", remoteFilePath, filePathRegex)}
	}
	bucket, err := receiver.client.Bucket(bucketName)
	if err != nil {
		return AliCloudError{err.Error()}
	}
	err = bucket.UploadFile(remoteFilePath, localFilePath, partSize)
	if err != nil {
		return AliCloudError{err.Error()}
	}
	return nil
}
