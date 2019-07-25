package responses

type GetTpAccessToken struct {
	Errno float64 `json:"errno"`
	Msg   string  `json:"msg"`
	Data  struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int64  `json:"expires_in"`
		Scope       string `json:"scope"`
	} `json:"data"`
}
