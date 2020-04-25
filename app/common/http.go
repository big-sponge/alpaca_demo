package common

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

/**
 * 封装访Request
 * @author ChengCheng
 * @date 2019-07-14 22:10:18
 */
func Request(input map[string]interface{}) (result interface{}, err error) {

	/*请求地址url*/
	reqUrl := input["url"].(string)
	method := "GET"
	if input["method"] != nil {
		method = input["method"].(string)
	}

	// 参数
	var bodyStr []byte
	if input["url_params"] != nil {
		urlStr := ""
		urlData := url.Values{}
		for key, val := range input["url_params"].(map[string]interface{}) {
			urlData.Add(key, val.(string))
		}
		urlStr = urlData.Encode()
		reqUrl += "?" + urlStr
	}
	if input["body_params"] != nil {
		bodyStr, err = json.Marshal(input["body_params"])
		if err != nil {
			return result, err
		}
	}

	// 请求
	request, err := http.NewRequest(method, reqUrl, strings.NewReader(string(bodyStr)))
	if err != nil {
		return result, err
	}

	/* Header信息 */
	if input["headers"] != nil {
		for key, val := range input["headers"].(map[string]interface{}) {
			request.Header.Set(key, val.(string))
		}
	} else {
		request.Header.Set("Content-Type", "application/json")
	}

	/* 超时时间 */
	timeOut := time.Second * 180
	if input["time_out"] != nil {
		timeOut = input["time_out"].(time.Duration)
	}
	client := &http.Client{
		Timeout: timeOut,
	}

	response, err := client.Do(request)
	if err != nil {
		return result, err
	}

	/* 获取body内容 */
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)

	/* 执行请求 */
	return result, err
}
