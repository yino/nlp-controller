package corp

type UserLoginResp struct {
	Token  string `json:"token"`
	Name   string `json:"name"`
	Mobile uint64 `json:"mobile"`
}
