package main

import (
	"fmt"
	"github.com/fadada-go-sdk-apiv3/api3req"
	"github.com/fadada-go-sdk-apiv3/bean"
	"github.com/fadada-go-sdk-apiv3/bean/account"
	"github.com/fadada-go-sdk-apiv3/client"
	"github.com/fadada-go-sdk-apiv3/test"
	"testing"
	"time"
)

/**
AppId:  "000003",
AppKey: "30BOgfz4vcEu7h6TjpYPa1EJ",
Url:    "http://127.0.0.1:8004/api/v3",
*/

var accountClient = client.AccountClient{
	Client: &client.ApiV3Client{AppId: "xxxx", AppKey: "xxxx",
		Url: "http://127.0.0.1:8004/api/v3",
		Req: &api3req.ApiV3Request{TimeOut: time.Duration(10) * time.Second}}}

/**
  获取个人unionId地址 请求示例
*/
func TestGetPersonUnionIdUrl(t *testing.T) {
	var req = account.GetPersonUnionIdUrlReq{
		ClientId: "12345",
		Notice: bean.Notice{
			NotifyWay:     1,
			NotifyAddress: "15013477347",
		},
		AllowModify: 1,
		RedirectUrl: "http://www.fadada.com",
		AuthScope:   "",
		AuthScheme:  1,
		Person: account.PersonReq{
			Name:                "xxxx",
			IdentType:           "",
			IdentNo:             "",
			Mobile:              "150xxxxxx",
			IdPhotoOptional:     0,
			BackIdCardImgBase64: "",
			IdCardImgBase64:     "",
			BankCardNo:          "",
			IsMiniProgram:       0,
		},
	}
	response, err := accountClient.GetPersonUnionIdUrl("76072b407d3447f1aa099f019f1f7b99", test.GetRandomString(32), req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(response)
}
