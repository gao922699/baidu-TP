package pages

import (
	"baidu-app/services"
	"baidu-app/structs/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 百度用户授权回调
func AuthRedirect(c *gin.Context) {
	clientId := c.Param("clientId")
	authorizationCode := c.Query("authorization_code")
	tpInfo, err := services.TpService{&models.Tp{}}.GetTpByClientId(clientId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	tokenInfo, err := services.AppService{}.GetAppAccessTokenByAuthorizationCode(tpInfo, authorizationCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	appInfo, err := services.AppService{}.GetAppInfo(tokenInfo.AccessToken)
	_ = services.AppService{}.CreateAppInfo(tpInfo, appInfo, tokenInfo)
	//成功后跳转到配置的回调地址
	html := "<p>授权成功，3秒后跳转。。。</p><script>setTimeout(\"location.href='" + tpInfo.AuthSuccessRedirectUrl + "'\",3000)</script>"
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(http.StatusOK, html)
}
