// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20220519

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type DescribeDataRepositoryTaskStatusRequestParams struct {
	// task id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// file system id
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

type DescribeDataRepositoryTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// task id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// file system id
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

func (r *DescribeDataRepositoryTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataRepositoryTaskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "FileSystemId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataRepositoryTaskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataRepositoryTaskStatusResponseParams struct {
	// ??????id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// ???????????? 0(????????????), 1(?????????), 2(?????????), 3(????????????)
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// ???????????? ID???????????????????????????????????????????????????????????????????????? RequestId???
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDataRepositoryTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataRepositoryTaskStatusResponseParams `json:"Response"`
}

func (r *DescribeDataRepositoryTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataRepositoryTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}