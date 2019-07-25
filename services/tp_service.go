package services

import (
	"baidu-app/databases"
	"baidu-app/structs/models"
	"baidu-app/structs/responses"
	"errors"
	"os"
	"time"
)

const url_auth = "https://smartprogram.baidu.com/mappconsole/tp/authorization" //用户授权跳转地址
const url_redrict = "/page/auth/redirect"                                      //用户授权后的回调地址

//小程序名称检测结果翻译
var checkResult = map[int]string{
	0: "检测通过",
	1: "重名检测未通过",
	2: "敏感&黄反检测未通过",
	3: "完全对等品牌词检测未通过",
	4: "包含品牌词检测未通过",
	5: "高流量词检测未通过"}

type TpService struct {
	Model *models.Tp
}

//根据clientId获取平台信息
func (tp TpService) GetTpByClientId(clientId string) (*models.Tp, error) {
	databases.Db.Where("client_id = ?", clientId).First(tp.Model)
	if tp.Model.ID == 0 {
		return tp.Model, errors.New("未找到平台信息")
	}
	return tp.Model, nil
}

//更新ticket
func (tp TpService) UpdateTicket(id int, ticket string) error {
	databases.Db.Model(&models.Tp{}).Where("id = ?", id).Update("ticket", ticket)
	return nil
}

//获取授权跳转地址
func (tp TpService) GetAuthUrl() (url string, err error) {
	accessToken, err := tp.getAccessToken()
	if err != nil {
		return "", errors.New(err.Error())
	}
	result, err := BaiduApi{}.GetPreAuthCode(accessToken)
	if err != nil {
		return "", errors.New("获取预授权码失败")
	}
	url = url_auth + "?client_id=" + tp.Model.ClientKey + "&redirect_uri=" + os.Getenv("PROJECT_DOMAIN") + url_redrict + "/" + tp.Model.ClientId + "&pre_auth_code=" + result.Data.PreAuthCode
	return url, nil
}

//检测小程序名称
func (tp TpService) CheckAppName(appName string) (result string, err error) {
	accessToken, err := tp.getAccessToken()
	if err != nil {
		return "", errors.New(err.Error())
	}
	response, err := BaiduApi{}.CheckAppName(accessToken, appName)
	if err != nil {
		return "", errors.New("检测失败")
	}
	result = checkResult[response.Data.CheckResult]
	return result, nil
}

//获取草稿列表
func (tp TpService) GetTemplateDraftList(page string, pageSize string) (result responses.GetTemplateDraftList, err error) {
	accessToken, err := tp.getAccessToken()
	if err != nil {
		return responses.GetTemplateDraftList{}, errors.New(err.Error())
	}
	response, err := BaiduApi{}.GetTemplateDraftList(accessToken, page, pageSize)
	if err != nil {
		return responses.GetTemplateDraftList{}, errors.New("检测失败")
	}
	return response, nil
}

//获取tp平台的access_token
func (tp TpService) getAccessToken() (accessToken string, err error) {
	if tp.Model.ID == 0 {
		return "", errors.New("未获取到tp信息")
	}
	//未过期直接返回数据库中的值
	if len(tp.Model.AccessToken) > 0 && ((tp.Model.ExpiredAt.Unix() - time.Now().Unix()) > 24*3600) {
		return tp.Model.AccessToken, nil
	} else {
		//已过期的重新请求接口获取
		result, err := BaiduApi{tp.Model.ClientId}.GetTpAccessToken(tp.Model.ClientKey, tp.Model.Ticket)
		if err != nil {
			return "", errors.New("获取access_token失败")
		}
		databases.Db.Model(&models.Tp{}).Updates(models.Tp{AccessToken: result.Data.AccessToken, ExpiredAt: time.Unix((time.Now().Unix() + int64(result.Data.ExpiresIn)), 0)})
		return result.Data.AccessToken, nil
	}
}
