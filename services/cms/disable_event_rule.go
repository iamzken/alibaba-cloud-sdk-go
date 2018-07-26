package cms

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

// DisableEventRule invokes the cms.DisableEventRule API synchronously
// api document: https://help.aliyun.com/api/cms/disableeventrule.html
func (client *Client) DisableEventRule(request *DisableEventRuleRequest) (response *DisableEventRuleResponse, err error) {
	response = CreateDisableEventRuleResponse()
	err = client.DoAction(request, response)
	return
}

// DisableEventRuleWithChan invokes the cms.DisableEventRule API asynchronously
// api document: https://help.aliyun.com/api/cms/disableeventrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DisableEventRuleWithChan(request *DisableEventRuleRequest) (<-chan *DisableEventRuleResponse, <-chan error) {
	responseChan := make(chan *DisableEventRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DisableEventRule(request)
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

// DisableEventRuleWithCallback invokes the cms.DisableEventRule API asynchronously
// api document: https://help.aliyun.com/api/cms/disableeventrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DisableEventRuleWithCallback(request *DisableEventRuleRequest, callback func(response *DisableEventRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DisableEventRuleResponse
		var err error
		defer close(result)
		response, err = client.DisableEventRule(request)
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

// DisableEventRuleRequest is the request struct for api DisableEventRule
type DisableEventRuleRequest struct {
	*requests.RpcRequest
	RuleNames *[]string `position:"Query" name:"RuleNames"  type:"Repeated"`
	RuleName  string    `position:"Query" name:"RuleName"`
}

// DisableEventRuleResponse is the response struct for api DisableEventRule
type DisableEventRuleResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDisableEventRuleRequest creates a request to invoke DisableEventRule API
func CreateDisableEventRuleRequest() (request *DisableEventRuleRequest) {
	request = &DisableEventRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2018-03-08", "DisableEventRule", "cms", "openAPI")
	return
}

// CreateDisableEventRuleResponse creates a response to parse from DisableEventRule response
func CreateDisableEventRuleResponse() (response *DisableEventRuleResponse) {
	response = &DisableEventRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}