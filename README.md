
## 法大大OpenApi3.0 go sdk demo
go version go1.13.1     

### 下载sdk


### 调用示例   
#### 获取token示例  
参考oauth2_test.go中的TestGetAccessToken测试方法

#### 获取个人unionId地址示例    
参考account_test.go中的TestPersonUnionIdUrl测试方法

#### 上传文件示例
参考file_test.go中的TestUploadFile测试方法

#### 下载签署文件示例   
参考file_test.go中的TestGetBySignFileIdReq测试方法

### 初始化client
以file_client.go中初始化fileClient示例，初始化需要appId，appKey， url，请求超时时间等参数
```go
var fileClient = client.FileClient{
	Client: &client.ApiV3Client{AppId: "000003", AppKey: "30BOgfz4vcEu7h6TjpYPa1EJ",
		Url: "http://127.0.0.1:8004/api/v3",
		Req: &api3req.ApiV3Request{TimeOut: time.Duration(10) * time.Second}}}
```
