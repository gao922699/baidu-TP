package apis

import (
	"baidu-app/services"
	"baidu-app/structs/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 检测小程序名称
func CheckAppName(c *gin.Context) {
	clientId := c.Param("clientId")
	appName := c.Param("name")
	tpInfo, err := services.TpService{&models.Tp{}}.GetTpByClientId(clientId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	result, err := services.TpService{tpInfo}.CheckAppName(appName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, result)
}
