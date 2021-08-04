package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"go.uber.org/zap"

	"github.com/yino/nlp-controller/config"
	"github.com/yino/nlp-controller/config/log"
	"github.com/yino/nlp-controller/domain/vo"
)

// QaApi api
type QaApi struct {
	Ak string
}

// MatchResp match resp
type MatchResp struct {
	Code    int64                      `json:"code"`
	Message string                     `json:"message"`
	Data    []vo.QaMatchQuestionItemVo `json:"data"`
}

// TrainModel 训练模型
func (a *QaApi) TrainModel() error {
	client := &http.Client{Timeout: 5 * time.Second}
	requestUrl := fmt.Sprintf("%s/qa/train_model?ak=%s", config.Conf.App.QaHost, a.Ak)
	resp, err := client.Get(requestUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	log.Logger.Error("main App Run error", zap.Error(err))
	apiResult, err := analysisResponseJson(resp)
	if err != nil {
		return err
	}
	if apiResult["code"] != float64(200) {
		return errors.New(apiResult["message"].(string))
	}

	return nil
}

// Match 检索问题
func (a *QaApi) Match(inputQuestion string) (result []vo.QaMatchQuestionItemVo, err error) {

	params := url.Values{}

	Url, err := url.Parse(fmt.Sprintf("%s/qa/match", config.Conf.App.QaHost))
	if err != nil {
		return
	}
	params.Set("ak", a.Ak)
	params.Set("question", inputQuestion)
	Url.RawQuery = params.Encode()
	client := &http.Client{Timeout: 5 * time.Second}
	requestUrl := Url.String()
	log.Logger.Info(requestUrl)
	resp, err := client.Get(requestUrl)
	if err != nil {
		log.Logger.Error("client error", zap.Error(err))
		return
	}
	defer resp.Body.Close()

	var body []byte
	var apiResult MatchResp
	body, err = ioutil.ReadAll(resp.Body)
	log.Logger.Info(string(body))
	if err != nil {
		log.Logger.Error("ioutil.ReadAll", zap.Error(err))
		return
	}
	err = json.Unmarshal(body, &apiResult)
	if err != nil {
		log.Logger.Error("json.Unmarshal", zap.Error(err))
		return
	}
	if apiResult.Code != 200 {
		log.Logger.Error("apiResult != 200", zap.Int64("code", apiResult.Code))
		return result, errors.New(apiResult.Message)
	}
	return apiResult.Data, nil
}

// analysisResponseJson 解析json
func analysisResponseJson(resp *http.Response) (map[string]interface{}, error) {
	var result map[string]interface{}
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		err = json.Unmarshal(body, &result)
	}
	return result, err
}
