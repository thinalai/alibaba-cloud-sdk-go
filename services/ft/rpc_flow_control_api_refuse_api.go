package ft

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

func (client *Client) RpcFlowControlApiRefuseApi(request *RpcFlowControlApiRefuseApiRequest) (response *RpcFlowControlApiRefuseApiResponse, err error) {
	response = CreateRpcFlowControlApiRefuseApiResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) RpcFlowControlApiRefuseApiWithChan(request *RpcFlowControlApiRefuseApiRequest) (<-chan *RpcFlowControlApiRefuseApiResponse, <-chan error) {
	responseChan := make(chan *RpcFlowControlApiRefuseApiResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RpcFlowControlApiRefuseApi(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) RpcFlowControlApiRefuseApiWithCallback(request *RpcFlowControlApiRefuseApiRequest, callback func(response *RpcFlowControlApiRefuseApiResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RpcFlowControlApiRefuseApiResponse
		var err error
		defer close(result)
		response, err = client.RpcFlowControlApiRefuseApi(request)
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

type RpcFlowControlApiRefuseApiRequest struct {
	*requests.RpcRequest
	RequiredValue     string `position:"Query" name:"RequiredValue"`
	HttpStatusCode    string `position:"Query" name:"HttpStatusCode"`
	Code              string `position:"Query" name:"Code"`
	IntValue          string `position:"Query" name:"IntValue"`
	StringValue       string `position:"Query" name:"StringValue"`
	EnumValue         string `position:"Query" name:"EnumValue"`
	Success           string `position:"Query" name:"Success"`
	DefaultValue      string `position:"Query" name:"DefaultValue"`
	NumberRange       string `position:"Query" name:"NumberRange"`
	SwitchValue       string `position:"Query" name:"SwitchValue"`
	Message           string `position:"Query" name:"Message"`
	ResultSwitchValue string `position:"Query" name:"ResultSwitchValue"`
}

type RpcFlowControlApiRefuseApiResponse struct {
	*responses.BaseResponse
	IntValue               string `json:"IntValue"`
	NumberRange            string `json:"NumberRange"`
	StringValue            string `json:"StringValue"`
	SwitchValue            string `json:"SwitchValue"`
	EnumValue              string `json:"EnumValue"`
	RequiredValue          string `json:"RequiredValue"`
	Success                string `json:"Success"`
	Code                   string `json:"Code"`
	Message                string `json:"Message"`
	HttpStatusCode         string `json:"HttpStatusCode"`
	NullToEmptyValue       string `json:"NullToEmptyValue"`
	ResultSwitchValue      string `json:"ResultSwitchValue"`
	NullToEmptyStructValue struct {
		NullToEmptyStructChildValue string `json:"NullToEmptyStructChildValue"`
	} `json:"NullToEmptyStructValue"`
	StructValue struct {
		StructChildValue string `json:"StructChildValue"`
	} `json:"StructValue"`
	ArrayValue []struct {
		ArrayChildValue string `json:"ArrayChildValue"`
	} `json:"ArrayValue"`
}

func CreateRpcFlowControlApiRefuseApiRequest() (request *RpcFlowControlApiRefuseApiRequest) {
	request = &RpcFlowControlApiRefuseApiRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ft", "2015-01-01", "RpcFlowControlApiRefuseApi", "", "")
	return
}

func CreateRpcFlowControlApiRefuseApiResponse() (response *RpcFlowControlApiRefuseApiResponse) {
	response = &RpcFlowControlApiRefuseApiResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
