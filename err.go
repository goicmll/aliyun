package aliyun

type AliCloudError struct {
	Msg string
}

func (receiver AliCloudError) Error() string {
	return receiver.Msg
}

func NewAliCloudError(Msg string) *AliCloudError {
	return &AliCloudError{Msg: Msg}
}
