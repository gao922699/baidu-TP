package responses

type CheckAppName struct {
	Errno float64 `json:"errno"`
	Msg   string  `json:"msg"`
	Data  struct {
		CheckResult int `json:"checkResult"`
	} `json:"data"`
}
