package apis

import (
	"baidu-app/services"
	"baidu-app/structs/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取小程序信息，并同步最新数据
func GetTemplateDraftList(c *gin.Context) {
	clientId := c.Param("clientId")
	page := c.DefaultQuery("page","1")
	pageSize := c.DefaultQuery("page_size","10")
	tpInfo, err := services.TpService{&models.Tp{}}.GetTpByClientId(clientId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	templateDraftList, err := services.TpService{tpInfo}.GetTemplateDraftList(page,pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, templateDraftList)
}
