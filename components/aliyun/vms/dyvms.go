package vms

import (
	"net/http"

	"github.com/denverdino/aliyungo/common"
)

// SendVmsArgs SendVmsArgs
type SendVmsArgs struct {
	CalledShowNumber string
	CalledNumber     string
	TtsCode          string
	TtsParam         string
	Volume           int
	PlayTimes        int
	OutId            string
}

// SendVmsResponse SendVmsResponse
type SendVmsResponse struct {
	common.Response
	Code    string
	Message string
	CallId  string
}

// DYVmsClient DYVmsClient
type DYVmsClient struct {
	common.Client
	Region common.Region
}

const (
	DYVmsEndPoint   = "https://dyvmsapi.aliyuncs.com/"
	SingleCallByTts = "SingleCallByTts"
	DYVmsAPIVersion = "2017-05-25"
)

// NewDYVmsClient NewDYVmsClient
func NewDYVmsClient(accessKeyId, accessKeySecret string) *DYVmsClient {
	client := new(DYVmsClient)
	client.Init(DYVmsEndPoint, DYVmsAPIVersion, accessKeyId, accessKeySecret)
	client.Region = common.Hangzhou
	return client
}

// SendVms SendVms
func (t *DYVmsClient) SendVms(args *SendVmsArgs) (*SendVmsResponse, error) {
	resp := SendVmsResponse{}
	return &resp, t.InvokeByAnyMethod(http.MethodGet, SingleCallByTts, "", args, &resp)
}
