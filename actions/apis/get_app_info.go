package apis

import (
	"baidu-app/services"
	"baidu-app/structs/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取小程序信息，并同步最新数据
func GetAppInfo(c *gin.Context) {
	appId := c.Param("appId")
	appModel, err := services.AppService{&models.App{}}.GetAppByAppId(appId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	accessToken, err := services.AppService{appModel}.GetAccessToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	appInfo, err := services.AppService{}.GetAppInfo(accessToken)
	//更新数据
	_ = services.AppService{}.UpdateAppInfo(appInfo)
	c.JSON(http.StatusOK, appInfo.Data)
}
