package responses

type GetPreAuthCode struct {
	Errno float64 `json:"errno"`
	Msg   string  `json:"msg"`
	Data  struct {
		PreAuthCode string `json:"pre_auth_code"`
		ExpiresIn   int64  `json:"expires_in"`
	} `json:"data"`
}
