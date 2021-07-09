package interfaces

// response code
const (
	//200 ~ 299 数据正常
	StatusSuccess = 200

	// 4000 ~ 4999
	ErrorUserNotFound = 4000
	ErrorNotLogin     = 4001

	// database err 5000~5999
	ErrorRegister    = 5001
	ErrorCreateToken = 5002

	// params error 6000~60001
	ErrorParams = 60001
)

var StatusText = map[int]string{
	StatusSuccess:     "success",
	ErrorUserNotFound: "user not fond",
	ErrorRegister:     "register fail",
	ErrorCreateToken:  "create token fail",
	ErrorParams:       "params error",
	ErrorNotLogin:     "please login first",
}
