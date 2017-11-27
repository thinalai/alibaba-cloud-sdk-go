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

func (client *Client) RpcHsfApi(request *RpcHsfApiRequest) (response *RpcHsfApiResponse, err error) {
	response = CreateRpcHsfApiResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) RpcHsfApiWithChan(request *RpcHsfApiRequest) (<-chan *RpcHsfApiResponse, <-chan error) {
	responseChan := make(chan *RpcHsfApiResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RpcHsfApi(request)
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

func (client *Client) RpcHsfApiWithCallback(request *RpcHsfApiRequest, callback func(response *RpcHsfApiResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RpcHsfApiResponse
		var err error
		defer close(result)
		response, err = client.RpcHsfApi(request)
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

type RpcHsfApiRequest struct {
	*requests.RpcRequest
	RequiredValue                  string `position:"Query" name:"RequiredValue"`
	Code                           string `position:"Query" name:"Code"`
	IntValue                       string `position:"Query" name:"IntValue"`
	DefaultValue                   string `position:"Query" name:"DefaultValue"`
	ProxyOriginalSecurityTransport string `position:"Query" name:"proxy_original_security_transport"`
	NumberRange                    string `position:"Query" name:"NumberRange"`
	Message                        string `position:"Query" name:"Message"`
	ResultSwitchValue              string `position:"Query" name:"ResultSwitchValue"`
	ProxyOriginalSourceIp          string `position:"Query" name:"proxy_original_source_ip"`
	JsonObject                     string `position:"Query" name:"JsonObject"`
	Sleep                          string `position:"Query" name:"Sleep"`
	HttpStatusCode                 string `position:"Query" name:"HttpStatusCode"`
	StringValue                    string `position:"Query" name:"StringValue"`
	RegionId                       string `position:"Query" name:"RegionId"`
	EnumValue                      string `position:"Query" name:"EnumValue"`
	Success                        string `position:"Query" name:"Success"`
	OtherParam                     string `position:"Query" name:"OtherParam"`
	SwitchValue                    string `position:"Query" name:"SwitchValue"`
	JsonObjectList                 string `position:"Query" name:"JsonObjectList"`
}

type RpcHsfApiResponse struct {
	*responses.BaseResponse
	IntValue                       string `json:"IntValue"`
	NumberRange                    string `json:"NumberRange"`
	StringValue                    string `json:"StringValue"`
	SwitchValue                    string `json:"SwitchValue"`
	EnumValue                      string `json:"EnumValue"`
	RequiredValue                  string `json:"RequiredValue"`
	Success                        string `json:"Success"`
	Code                           string `json:"Code"`
	Message                        string `json:"Message"`
	HttpStatusCode                 string `json:"HttpStatusCode"`
	NullToEmptyValue               string `json:"NullToEmptyValue"`
	ResultSwitchValue              string `json:"ResultSwitchValue"`
	RegionId                       string `json:"RegionId"`
	Name                           string `json:"Name"`
	Age                            string `json:"Age"`
	JsonListStrig                  string `json:"JsonListStrig"`
	CallerUid                      string `json:"callerUid"`
	CallerBid                      string `json:"callerBid"`
	ProxyTrustTransportInfo        string `json:"proxy_trust_transport_info"`
	AkMfaPresent                   string `json:"ak_mfa_present"`
	CallerType                     string `json:"callerType"`
	CallerParentId                 string `json:"callerParentId"`
	ResourceOwnerId                string `json:"resourceOwnerId"`
	AppIp                          string `json:"app_ip"`
	ProxyOriginalSourceIp          string `json:"proxy_original_source_ip"`
	ProxyOriginalSecurityTransport string `json:"proxy_original_security_transport"`
	SecurityTransport              string `json:"security_transport"`
	RequestId                      string `json:"requestId"`
	NullToEmptyStructValue         struct {
		NullToEmptyStructChildValue string `json:"NullToEmptyStructChildValue"`
	} `json:"NullToEmptyStructValue"`
	StructValue struct {
		StructChildValue string `json:"StructChildValue"`
	} `json:"StructValue"`
	ArrayValue []struct {
		ArrayChildValue string `json:"ArrayChildValue"`
	} `json:"ArrayValue"`
}

func CreateRpcHsfApiRequest() (request *RpcHsfApiRequest) {
	request = &RpcHsfApiRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ft", "2015-01-01", "RpcHsfApi", "", "")
	return
}

func CreateRpcHsfApiResponse() (response *RpcHsfApiResponse) {
	response = &RpcHsfApiResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
