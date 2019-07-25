package responses

type GetTemplateDraftList struct {
	Errno float64 `json:"errno"`
	Msg   string  `json:"msg"`
	Data  interface{}
}