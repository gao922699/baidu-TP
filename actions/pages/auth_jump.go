package pages

import (
	"baidu-app/services"
	"baidu-app/structs/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 跳转百度授权页面
func AuthJump(c *gin.Context) {
	clientId := c.Param("clientId")
	tpInfo, err := services.TpService{&models.Tp{}}.GetTpByClientId(clientId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	jumpUrl,err := services.TpService{tpInfo}.GetAuthUrl()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	html := "<p>跳转中，请稍侯。。。</p><script>setTimeout(\"location.href='" + jumpUrl + "'\",500)</script>"
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(http.StatusOK, html)
}
