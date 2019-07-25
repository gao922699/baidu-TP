## 百度小程序平台
用于百度TP帐号管理小程序的平台
-	POST("/callback/auth/:clientId") 接收百度推送（获取ticket）
-	GET("/page/auth/jump/:clientId") 跳转到授权页面
-	GET("/page/auth/redirect/:clientId") 授权页面回调地址
-	GET("/app/:appId") 获取小程序详细信息
-	GET("tp/:clientId/app/check-name/:name") 检测小程序名称
-	GET("tp/:clientId/template-draft-list") 模板草稿列表

## 使用
- 拷贝 .env.sample 到 .env,修改相关参数
- go run main.go