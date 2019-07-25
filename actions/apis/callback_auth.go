package apis

import (
	"baidu-app/services"
	"baidu-app/structs/models"
	"baidu-app/structs/requests"
	"baidu-app/structs/responses"
	"baidu-app/util"
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 接收百度TP的服务器推送
func ReceiveMessage(c *gin.Context) {
	clientId := c.Param("clientId")
	params := requests.CallbackAuth{}
	_ = c.BindJSON(&params)
	tpInfo, err := services.TpService{&models.Tp{}}.GetTpByClientId(clientId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	//base64_decode接收到的encrypt
	encrypt, _ := base64.StdEncoding.DecodeString(params.Encrypt)
	//用解码方法解码encrypt
	decodeUtil := util.NewBaiduEncrypt(tpInfo.DecodeKey, tpInfo.ClientId)
	decodedEncrypt, err := decodeUtil.Decode(encrypt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	//更新ticket
	decodedStruct := responses.CallbackAuth{}
	_ = json.Unmarshal([]byte(decodedEncrypt), &decodedStruct)
	_ = services.TpService{}.UpdateTicket(tpInfo.ID, decodedStruct.Ticket)

	c.String(http.StatusOK, "success")
}
