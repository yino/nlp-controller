package interfaces

// response code
const (
	//200 ~ 299 数据正常
	StatusSuccess = 200

	// 4000 ~ 4999
	ErrorUserNotFound = 4000
	ErrorNotLogin     = 4001
	ErrorToken        = 4002

	// database err 5000~5999
	ErrorRegister     = 5001
	ErrorCreateToken  = 5002
	ErrorGetData      = 5003
	ErrorCreateData   = 5004
	ErrorDataNoteUser = 5005
	ErrorUpdateData   = 5006
	ErrorDeleteData   = 5007

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
	ErrorToken:        "invalid token",
	ErrorGetData:      "get data fail",
	ErrorCreateData:   "create data fail",
	ErrorDataNoteUser: "This data does not belong to the user",
	ErrorUpdateData:   "update data fail",
	ErrorDeleteData:   "delete data fail",
}
