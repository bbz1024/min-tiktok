package transform

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"strings"
	"time"
)

type NlsTask struct {
	client       *sdk.Client
	regionId     string
	endpointName string
	domain       string
	product      string
	apiVersion   string
	postAction   string
	getAction    string
	accessKeyId  string
	accessKeySec string
	appKey       string
}
type NlsResponse struct {
	StatusText string
	StatusCode int
	TaskId     string
	Content    string
}

func NewNlsTask(accessKeyId, accessKeySec, appKey string) *NlsTask {
	return &NlsTask{
		regionId:     "cn-shanghai",
		endpointName: "cn-shanghai",
		domain:       "filetrans.cn-shanghai.aliyuncs.com",
		product:      "nls-filetrans",
		apiVersion:   "2018-08-17",
		postAction:   "SubmitTask",
		getAction:    "GetTaskResult",
		accessKeyId:  accessKeyId,
		accessKeySec: accessKeySec,
		appKey:       appKey,
	}
}

func (nt *NlsTask) SubmitTask(fileLink string) (string, error) {
	client, err := sdk.NewClientWithAccessKey(nt.regionId, nt.accessKeyId, nt.accessKeySec)
	if err != nil {
		return "", err
	}
	nt.client = client

	postRequest := requests.NewCommonRequest()
	postRequest.Domain = nt.domain
	postRequest.Version = nt.apiVersion
	postRequest.Product = nt.product
	postRequest.ApiName = nt.postAction
	postRequest.Method = "POST"

	mapTask := map[string]string{
		"appkey":                      nt.appKey,
		"file_link":                   fileLink,
		"version":                     "4.0",
		"enable_words":                "false",
		"enable_sample_rate_adaptive": "true",
		"enable_disfluency":           "true",
	}

	taskJson, err := json.Marshal(mapTask)
	if err != nil {
		return "", err
	}

	postRequest.FormParams["Task"] = string(taskJson)

	postResponse, err := nt.client.ProcessCommonRequest(postRequest)
	if err != nil {
		return "", err
	}

	if postResponse.GetHttpStatus() != 200 {
		return "", fmt.Errorf("录音文件识别请求失败，Http错误码: %d", postResponse.GetHttpStatus())
	}

	var postMapResult map[string]interface{}
	if err := json.Unmarshal([]byte(postResponse.GetHttpContentString()), &postMapResult); err != nil {
		return "", err
	}
	fmt.Println(postMapResult)
	return postMapResult["TaskId"].(string), nil
}

func (nt *NlsTask) GetTaskResult(taskID string) (*NlsResponse, error) {
	if nt.client == nil {
		return nil, fmt.Errorf("client not initialized")
	}

	getRequest := requests.NewCommonRequest()
	getRequest.Domain = nt.domain
	getRequest.Version = nt.apiVersion
	getRequest.Product = nt.product
	getRequest.ApiName = nt.getAction
	getRequest.Method = "GET"
	getRequest.QueryParams["TaskId"] = taskID

	for {
		getResponse, err := nt.client.ProcessCommonRequest(getRequest)
		if err != nil {
			return nil, err
		}

		if getResponse.GetHttpStatus() != 200 {
			return nil, fmt.Errorf("识别结果查询请求失败，Http错误码：%d", getResponse.GetHttpStatus())
		}

		var getMapResult map[string]interface{}
		if err := json.Unmarshal([]byte(getResponse.GetHttpContentString()), &getMapResult); err != nil {
			return nil, err
		}
		statusText := getMapResult["StatusText"].(string)
		if statusText == "SUCCESS" {
			var content strings.Builder
			/*
				map[
				BizDuration:3101
				RequestId:0D8CB910-E1EF-58C9-B53A-9B3F567F88D5
				RequestTime:1.720481362987e+12
				Result:
					map[Sentences:
						[
						map[BeginTime:820 ChannelId:0 EmotionValue:6.6 EndTime:3080
								SilenceDuration:0 SpeakerId:1 SpeechRate:132 Text:北京的天气。]
						]
					]
				SolveTime:1.720481363502e+12
				StatusCode:2.105e+07
				StatusText:SUCCESS
				TaskId:645b7b715da248bea6dd95bafed642fb
				]
			*/
			rest := getMapResult["Result"]
			mapResult := rest.(map[string]interface{})
			sentences := mapResult["Sentences"].([]interface{})
			for _, sentence := range sentences {
				mapSentence := sentence.(map[string]interface{})
				content.WriteString(mapSentence["Text"].(string))
			}
			return &NlsResponse{
				StatusText: statusText,
				StatusCode: int(getMapResult["StatusCode"].(float64)),
				TaskId:     taskID,
				Content:    content.String(),
			}, nil
		} else if statusText != "RUNNING" && statusText != "QUEUEING" {
			return nil, fmt.Errorf("任务状态异常: %s", statusText)
		}
		time.Sleep(10 * time.Second)
	}
}
