package responses

type GetAppInfo struct {
	Errno float64 `json:"errno"`
	Msg   string  `json:"msg"`
	Data  struct {
		AppId         float64 `json:"app_id"`
		AppKey        string  `json:"app_key"`
		AppName       string  `json:"app_name"`
		AppDesc       string  `json:"app_desc"`
		PhotoAddr     string  `json:"photo_addr"`
		Status        int     `json:"status"`
		WebStatus     int     `json:"web_status"`
		Qualification struct {
			Name     string `json:"name"`
			Type     int    `json:"type"`
			Status   int    `json:"status"`
			AdType   int    `json:"ad_type"`
			AdStatus int    `json:"ad_status"`
		} `json:"qualification"`
		ModifyCount struct {
			CategoryModifyQuota  int `json:"category_modify_quota"`
			CategoryModifyUsed   int `json:"category_modify_used"`
			ImageModifyQuota     int `json:"image_modify_quota"`
			ImageModifyUsed      int `json:"image_modify_used"`
			NameModifyQuota      int `json:"name_modify_quota"`
			NameModifyUsed       int `json:"name_modify_used"`
			SignatureModifyQuota int `json:"signature_modify_quota"`
			SignatureModifyUsed  int `json:"signature_modify_used"`
		}
		MinSwanVersion string `json:"min_swan_version"`
		Category       struct {
			AuditStatus  int     `json:"audit_status"`
			CategoryDesc string  `json:"category_desc"`
			CategoryId   float64 `json:"category_id"`
			CategoryName string  `json:"category_name"`
			Parent       struct {
				CategoryId   float64 `json:"category_id"`
				CategoryDesc string  `json:"category_desc"`
				CategoryName string  `json:"category_name"`
			} `json:"parent"`
		} `json:"category"`
		AuthInfo  interface{} `json:"auth_info"`
		AuditInfo struct {
			AuditAppName       string `json:"audit_app_name"`
			AuditAppNameReason string `json:"audit_app_name_reason"`
			AuditAppNameStatus int    `json:"audit_app_name_status"`
		} `json:"audit_info"`
	} `json:"data"`
}
