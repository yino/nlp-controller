package interfaces

type BaseResp struct {
	Code int64 `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}
