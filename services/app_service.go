package services

import (
	"baidu-app/databases"
	"baidu-app/structs/models"
	"baidu-app/structs/responses"
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

type AppService struct {
	Model *models.App
}

//根据appId获取app信息
func (app AppService) GetAppByAppId(appId string) (*models.App, error) {
	databases.Db.Where("app_id = ?", appId).Preload("Tp").First(app.Model)
	if app.Model.ID == 0 {
		return app.Model, errors.New("未找到APP信息")
	}
	return app.Model, nil
}

//根据authCode获取app的accessToken
func (app AppService) GetAppAccessTokenByAuthorizationCode(tp *models.Tp, authorizationCode string) (result responses.GetAppAccessToken, err error) {
	result, err = BaiduApi{tp.ClientId}.GetAppAccessTokenByAuthorizationCode(tp.AccessToken, authorizationCode)
	if err != nil {
		return responses.GetAppAccessToken{}, errors.New("获取access_token失败")
	}
	return result, nil
}

//创建一条APP记录
func (app AppService) CreateAppInfo(tpInfo *models.Tp, appInfo responses.GetAppInfo, tokenInfo responses.GetAppAccessToken) error {
	appId := strconv.FormatFloat(appInfo.Data.AppId, 'E', -1, 64)
	appInfoJson, _ := json.Marshal(appInfo.Data)
	databases.Db.Create(models.App{
		TpId:         tpInfo.ID,
		AppId:        appId,
		AppName:      appInfo.Data.AppName,
		Data:         string(appInfoJson),
		AccessToken:  tokenInfo.AccessToken,
		RefreshToken: tokenInfo.RefreshToken,
		ExpiredAt:    time.Unix((time.Now().Unix() + int64(tokenInfo.ExpiresIn)), 0),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now()})
	return nil
}

//获取accessToken
func (app AppService) GetAccessToken() (accessToken string, err error) {
	if app.Model.ID == 0 {
		return "", errors.New("未获取到app信息")
	}
	if len(app.Model.AccessToken) > 0 && ((app.Model.ExpiredAt.Unix() - time.Now().Unix()) > 24*3600) {
		return app.Model.AccessToken, nil
	} else {
		tokenInfo, err := BaiduApi{}.GetAppAccessToken(app.Model.Tp.AccessToken, app.Model.RefreshToken)
		if err != nil {
			return "", errors.New("获取accessToken失败")
		}
		databases.Db.Model(app.Model).Updates(models.App{
			AccessToken:  tokenInfo.AccessToken,
			RefreshToken: tokenInfo.RefreshToken,
			ExpiredAt:    time.Unix((time.Now().Unix() + int64(tokenInfo.ExpiresIn)), 0)})
		return tokenInfo.AccessToken, nil
	}
}

//获取app详情
func (app AppService) GetAppInfo(accessToken string) (result responses.GetAppInfo, err error) {
	result, err = BaiduApi{}.GetAppInfo(accessToken)
	return result, err
}

//更新APP信息
func (app AppService) UpdateAppInfo(appInfo responses.GetAppInfo) error {
	jsonAppInfo, _ := json.Marshal(appInfo.Data)
	databases.Db.Model(&models.App{}).Where("app_id = ?", appInfo.Data.AppId).Updates(models.App{
		AppName: appInfo.Data.AppName,
		Data:    string(jsonAppInfo)})
	return nil
}
