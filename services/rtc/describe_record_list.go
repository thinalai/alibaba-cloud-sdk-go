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

package rtc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeRecordList invokes the rtc.DescribeRecordList API synchronously
// api document: https://help.aliyun.com/api/rtc/describerecordlist.html
func (client *Client) DescribeRecordList(request *DescribeRecordListRequest) (response *DescribeRecordListResponse, err error) {
	response = CreateDescribeRecordListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRecordListWithChan invokes the rtc.DescribeRecordList API asynchronously
// api document: https://help.aliyun.com/api/rtc/describerecordlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRecordListWithChan(request *DescribeRecordListRequest) (<-chan *DescribeRecordListResponse, <-chan error) {
	responseChan := make(chan *DescribeRecordListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRecordList(request)
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

// DescribeRecordListWithCallback invokes the rtc.DescribeRecordList API asynchronously
// api document: https://help.aliyun.com/api/rtc/describerecordlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRecordListWithCallback(request *DescribeRecordListRequest, callback func(response *DescribeRecordListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRecordListResponse
		var err error
		defer close(result)
		response, err = client.DescribeRecordList(request)
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

// DescribeRecordListRequest is the request struct for api DescribeRecordList
type DescribeRecordListRequest struct {
	*requests.RpcRequest
	OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
	AppId       string           `position:"Query" name:"AppId"`
	StartTime   string           `position:"Query" name:"StartTime"`
	EndTime     string           `position:"Query" name:"EndTime"`
	IdType      string           `position:"Query" name:"IdType"`
	Id          string           `position:"Query" name:"Id"`
	SortType    string           `position:"Query" name:"SortType"`
	ServiceArea string           `position:"Query" name:"ServiceArea"`
	PageNo      requests.Integer `position:"Query" name:"PageNo"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
}

// DescribeRecordListResponse is the response struct for api DescribeRecordList
type DescribeRecordListResponse struct {
	*responses.BaseResponse
	RequestId                string                                      `json:"RequestId" xml:"RequestId"`
	PageSize                 int64                                       `json:"PageSize" xml:"PageSize"`
	PageNo                   int64                                       `json:"PageNo" xml:"PageNo"`
	TotalCnt                 int64                                       `json:"TotalCnt" xml:"TotalCnt"`
	CommunicationRecordInfos DescribeRecordListCommunicationRecordInfos0 `json:"CommunicationRecordInfos" xml:"CommunicationRecordInfos"`
}

type DescribeRecordListCommunicationRecordInfos0 struct {
	CommunicationRecordInfo []DescribeRecordListCommunicationRecordInfo1 `json:"CommunicationRecordInfo" xml:"CommunicationRecordInfo"`
}

type DescribeRecordListCommunicationRecordInfo1 struct {
	ChannelId    string   `json:"ChannelId" xml:"ChannelId"`
	StartTime    string   `json:"StartTime" xml:"StartTime"`
	EndTime      string   `json:"EndTime" xml:"EndTime"`
	TotalUserCnt int64    `json:"TotalUserCnt" xml:"TotalUserCnt"`
	Status       bool     `json:"Status" xml:"Status"`
	RecordId     string   `json:"RecordId" xml:"RecordId"`
	CallAreas    []string `json:"CallAreas" xml:"CallAreas"`
}

// CreateDescribeRecordListRequest creates a request to invoke DescribeRecordList API
func CreateDescribeRecordListRequest() (request *DescribeRecordListRequest) {
	request = &DescribeRecordListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("rtc", "2018-01-11", "DescribeRecordList", "rtc", "openAPI")
	return
}

// CreateDescribeRecordListResponse creates a response to parse from DescribeRecordList response
func CreateDescribeRecordListResponse() (response *DescribeRecordListResponse) {
	response = &DescribeRecordListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
