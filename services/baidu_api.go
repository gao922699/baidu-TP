package services

import (
	"baidu-app/structs/responses"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const get_tp_access_token = "https://openapi.baidu.com/public/2.0/smartapp/auth/tp/token"                   //用ticket换取TP平台的access_token
const get_pre_auth_code = "https://openapi.baidu.com/rest/2.0/smartapp/tp/createpreauthcode"                //获取预授权码
const get_app_access_token = "https://openapi.baidu.com/rest/2.0/oauth/token"                               //获取小程序的access_token
const get_app_info = "https://openapi.baidu.com/rest/2.0/smartapp/app/info"                                 //获取已授权的小程序的详情
const get_check_app_name = "https://openapi.baidu.com/rest/2.0/smartapp/app/checkname"                      //小程序名称检测
const get_template_draft_list = "https://openapi.baidu.com/rest/2.0/smartapp/template/gettemplatedraftlist" //获取TP账号下的模板草稿列表

type BaiduApi struct {
	ClientId string
}

//获取tp平台的access_token
func (bd BaiduApi) GetTpAccessToken(clientKey string, ticket string) (result responses.GetTpAccessToken, err error) {
	url := get_tp_access_token + "?client_id=" + clientKey + "&ticket=" + ticket
	response, err := get(url)
	err = json.Unmarshal(response, &result)
	if result.Errno != 0 {
		return result, errors.New("百度接口错误：" + result.Msg)
	}
	return result, err
}

//获取预授权码
func (bd BaiduApi) GetPreAuthCode(tpAccessToken string) (result responses.GetPreAuthCode, err error) {
	url := get_pre_auth_code + "?access_token=" + tpAccessToken
	response, err := get(url)
	err = json.Unmarshal(response, &result)
	if result.Errno != 0 {
		return result, errors.New("百度接口错误：" + result.Msg)
	}
	return result, err
}

//根据authcode获取accesstoken
func (bd BaiduApi) GetAppAccessTokenByAuthorizationCode(tpAccessToken string, authorizationCode string) (result responses.GetAppAccessToken, err error) {
	url := get_app_access_token + "?access_token=" + tpAccessToken + "&code=" + authorizationCode + "&grant_type=app_to_tp_authorization_code"
	response, err := get(url)
	err = json.Unmarshal(response, &result)
	if len(result.AccessToken) == 0 {
		return result, errors.New("百度接口错误")
	}
	return result, err
}

//获取（刷新）小程序access_token
func (bd BaiduApi) GetAppAccessToken(tpAccessToken string, refreshToken string) (result responses.GetAppAccessToken, err error) {
	url := get_app_access_token + "?access_token=" + tpAccessToken + "&refresh_token=" + refreshToken + "&grant_type=app_to_tp_refresh_token"
	response, err := get(url)
	err = json.Unmarshal(response, &result)
	if len(result.AccessToken) == 0 {
		return result, errors.New("百度接口错误")
	}
	return result, err
}

//获取小程序详细信息
func (bd BaiduApi) GetAppInfo(accessToken string) (result responses.GetAppInfo, err error) {
	url := get_app_info + "?access_token=" + accessToken
	response, err := get(url)
	err = json.Unmarshal(response, &result)
	if result.Errno != 0 {
		return result, errors.New("百度接口错误：" + result.Msg)
	}
	return result, err
}

//检测小程序名称
func (bd BaiduApi) CheckAppName(tpAccessToken string, appName string) (result responses.CheckAppName, err error) {
	url := get_check_app_name + "?access_token=" + tpAccessToken + "&app_name=" + appName
	response, err := get(url)
	err = json.Unmarshal(response, &result)
	if result.Errno != 0 {
		return result, errors.New("百度接口错误：" + result.Msg)
	}
	return result, err
}

//获取TP平台模板草稿列表
func (bd BaiduApi) GetTemplateDraftList(tpAccessToken string, page string, pageSize string) (result responses.GetTemplateDraftList, err error) {
	url := get_template_draft_list + "?access_token=" + tpAccessToken + "&page" + page + "&page_size" + pageSize
	response, err := get(url)
	err = json.Unmarshal(response, &result)
	if result.Errno != 0 {
		return result, errors.New("百度接口错误：" + result.Msg)
	}
	return result, err
}

//发送get请求
func get(url string) (result []byte, err error) {
	response, err := http.Get(url)
	if err != nil {
		return []byte{}, errors.New("百度接口访问失败")
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	//请求记录日志
	file, _ := os.OpenFile("logs/request_log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer file.Close()
	log.SetOutput(file)
	log.SetPrefix("[Info]")
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)
	log.Println(url)
	log.Println(string(body))
	return body, nil
}
