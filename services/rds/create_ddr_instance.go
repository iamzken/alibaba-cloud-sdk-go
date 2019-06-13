package rds

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// CreateDdrInstance invokes the rds.CreateDdrInstance API synchronously
// api document: https://help.aliyun.com/api/rds/createddrinstance.html
func (client *Client) CreateDdrInstance(request *CreateDdrInstanceRequest) (response *CreateDdrInstanceResponse, err error) {
	response = CreateCreateDdrInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDdrInstanceWithChan invokes the rds.CreateDdrInstance API asynchronously
// api document: https://help.aliyun.com/api/rds/createddrinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateDdrInstanceWithChan(request *CreateDdrInstanceRequest) (<-chan *CreateDdrInstanceResponse, <-chan error) {
	responseChan := make(chan *CreateDdrInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDdrInstance(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// CreateDdrInstanceWithCallback invokes the rds.CreateDdrInstance API asynchronously
// api document: https://help.aliyun.com/api/rds/createddrinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateDdrInstanceWithCallback(request *CreateDdrInstanceRequest, callback func(response *CreateDdrInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDdrInstanceResponse
		var err error
		defer close(result)
		response, err = client.CreateDdrInstance(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// CreateDdrInstanceRequest is the request struct for api CreateDdrInstance
type CreateDdrInstanceRequest struct {
	*requests.RpcRequest
	ConnectionMode        string           `position:"Query" name:"ConnectionMode"`
	ResourceOwnerId       requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBInstanceStorage     requests.Integer `position:"Query" name:"DBInstanceStorage"`
	SystemDBCharset       string           `position:"Query" name:"SystemDBCharset"`
	SourceDBInstanceName  string           `position:"Query" name:"SourceDBInstanceName"`
	ClientToken           string           `position:"Query" name:"ClientToken"`
	HostType              string           `position:"Query" name:"HostType"`
	EngineVersion         string           `position:"Query" name:"EngineVersion"`
	UserBakSetURL         string           `position:"Query" name:"UserBakSetURL"`
	ResourceGroupId       string           `position:"Query" name:"ResourceGroupId"`
	Engine                string           `position:"Query" name:"Engine"`
	DBInstanceDescription string           `position:"Query" name:"DBInstanceDescription"`
	DBInstanceStorageType string           `position:"Query" name:"DBInstanceStorageType"`
	BackupSetRegion       string           `position:"Query" name:"BackupSetRegion"`
	DBInstanceNetType     string           `position:"Query" name:"DBInstanceNetType"`
	BackupSetType         string           `position:"Query" name:"BackupSetType"`
	Period                string           `position:"Query" name:"Period"`
	RestoreTime           string           `position:"Query" name:"RestoreTime"`
	BakSetName            string           `position:"Query" name:"BakSetName"`
	ResourceOwnerAccount  string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string           `position:"Query" name:"OwnerAccount"`
	BackupSetId           string           `position:"Query" name:"BackupSetId"`
	OwnerId               requests.Integer `position:"Query" name:"OwnerId"`
	UsedTime              string           `position:"Query" name:"UsedTime"`
	DBInstanceClass       string           `position:"Query" name:"DBInstanceClass"`
	SecurityIPList        string           `position:"Query" name:"SecurityIPList"`
	VSwitchId             string           `position:"Query" name:"VSwitchId"`
	PrivateIpAddress      string           `position:"Query" name:"PrivateIpAddress"`
	RestoreType           string           `position:"Query" name:"RestoreType"`
	VPCId                 string           `position:"Query" name:"VPCId"`
	TunnelId              string           `position:"Query" name:"TunnelId"`
	ZoneId                string           `position:"Query" name:"ZoneId"`
	PayType               string           `position:"Query" name:"PayType"`
	SourceRegion          string           `position:"Query" name:"SourceRegion"`
	InstanceNetworkType   string           `position:"Query" name:"InstanceNetworkType"`
}

// CreateDdrInstanceResponse is the response struct for api CreateDdrInstance
type CreateDdrInstanceResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	DBInstanceId     string `json:"DBInstanceId" xml:"DBInstanceId"`
	OrderId          string `json:"OrderId" xml:"OrderId"`
	ConnectionString string `json:"ConnectionString" xml:"ConnectionString"`
	Port             string `json:"Port" xml:"Port"`
}

// CreateCreateDdrInstanceRequest creates a request to invoke CreateDdrInstance API
func CreateCreateDdrInstanceRequest() (request *CreateDdrInstanceRequest) {
	request = &CreateDdrInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "CreateDdrInstance", "rds", "openAPI")
	return
}

// CreateCreateDdrInstanceResponse creates a response to parse from CreateDdrInstance response
func CreateCreateDdrInstanceResponse() (response *CreateDdrInstanceResponse) {
	response = &CreateDdrInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
