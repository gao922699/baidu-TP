package main

import (
	"baidu-app/actions/apis"
	"baidu-app/actions/pages"
	"baidu-app/databases"
	"baidu-app/structs/models"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

func main() {
	databases.Db.Set("gorm:table_options", "ENGINE=InnoDB,CHARSET=utf8").AutoMigrate(&models.Tp{}, &models.App{})
	//初始化一条tp数据
	//databases.Db.Where(models.Tp{Name: "放心投", ClientId: "16520930", ClientKey: "lHEmtBDGVKkL1j7THzfmOvEDksCr9ZKy", DecodeKey: "3awOAp37nrne7Lc8JOBnxB0znbsnZSdCcYMi5A5hqs6"}).FirstOrCreate(&models.Tp{})
	clientNames := strings.Split(os.Getenv("TP_INFO_CLIENT_NAMES"),",")
	clientIds := strings.Split(os.Getenv("TP_INFO_CLIENT_IDS"),",")
	clientKeys := strings.Split(os.Getenv("TP_INFO_CLIENT_KEYS"),",")
	clientDecodeKeys := strings.Split(os.Getenv("TP_INFO_CLIENT_DECODE_KEYS"),",")
	clientAuthSuccessRedirectUrls := strings.Split(os.Getenv("TP_INFO_CLIENT_AUTH_SUCCESS_REDIRECT_URLS"),",")
	for i:=0;i<len(clientNames);i++ {
		databases.Db.Where(models.Tp{
			Name: clientNames[i],
			ClientId: clientIds[i],
			ClientKey: clientKeys[i],
			DecodeKey: clientDecodeKeys[i],
			AuthSuccessRedirectUrl: clientAuthSuccessRedirectUrls[i],
		}).FirstOrCreate(&models.Tp{})
	}
	router := gin.Default()

	router.POST("/callback/auth/:clientId", apis.ReceiveMessage)
	router.GET("/page/auth/jump/:clientId", pages.AuthJump)
	router.GET("/page/auth/redirect/:clientId", pages.AuthRedirect)
	router.GET("/app/:appId", apis.GetAppInfo)
	router.GET("tp/:clientId/app/check-name/:name", apis.CheckAppName)
	router.GET("tp/:clientId/template-draft-list", apis.GetTemplateDraftList)

	router.Run(":8080")
}
