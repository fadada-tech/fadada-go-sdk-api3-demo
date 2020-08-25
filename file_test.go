package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"github.com/axgle/mahonia"
	"github.com/fadada-go-sdk-apiv3/api3req"
	"github.com/fadada-go-sdk-apiv3/bean/file"
	"github.com/fadada-go-sdk-apiv3/client"
	"github.com/fadada-go-sdk-apiv3/test"
	"io"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

//初始化ApiV3Client，默认使用
var fileClient = client.FileClient{
	Client: &client.ApiV3Client{AppId: "xxxx", AppKey: "xxxx",
		Url: "xxxx",
		Req: &api3req.ApiV3Request{TimeOut: time.Duration(10) * time.Second}}}

/**
  上传文件 请求示例
*/
func TestUploadFile(t *testing.T) {
	uploadFile, err := os.Open("C:/Users/huj1/Desktop/temp/授权模板.pdf")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file hash = {}", api3req.GetFileHash(uploadFile))
	var req file.UploadFileReq
	req.FileType = 1
	response, err := fileClient.UploadFile("9bdf3fc5b18a4b3fbb214770c8371c29", test.GetRandomString(32), req, uploadFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(response)
}

/**
  下载签署文件 请求示例
*/
func TestGetBySignFileIdReq(t *testing.T) {
	var req file.GetBySignFileIdReq
	req.SignFileId = ""
	req.TaskId = "126e9510b66e47068a4b92a16bd84331"
	response, err := fileClient.GetBySignFileId("b63eaa18eae84bfe844e214ab1e593df", test.GetRandomString(32), req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("ContentType = {}", response.ContentType)
	if response.ContentType == "application/zip" {
		//解压zip示例
		z, err := zip.NewReader(bytes.NewReader(response.Bytes), int64(len(response.Bytes)))
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, zf := range z.File {
			fnm := mahonia.NewDecoder("gbk").ConvertString(zf.Name)
			fmt.Println(fnm)
			f, err := zf.Open()
			if err != nil {
				fmt.Println(err)
				return
			}
			buf := &bytes.Buffer{}
			if _, err := io.Copy(buf, f); err != nil {
				fmt.Println(err)
				return
			}
			if err := ioutil.WriteFile("C:/Users/huj1/Desktop/temp/"+fnm, buf.Bytes(), 0655); err != nil {
				fmt.Println(err)
			}
		}
		/*if err := ioutil.WriteFile("C:/Users/huj1/Desktop/temp/126e9510b66e47068a4b92a16bd84331.zip", response.Bytes, 0655); err != nil {
			fmt.Println(err)
		}*/
	} else {
		fmt.Println("response = {}", string(response.Bytes))
	}
}
