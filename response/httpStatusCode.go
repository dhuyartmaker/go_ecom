package response

const (
	ErrCodeSuccess      = 20001
	ErrCodeParamInvalid = 20003
	ErrorCodeAuthorization = 40001
)

var Msg = map[int]string{
	ErrCodeSuccess:      "success",
	ErrCodeParamInvalid: "Invalid",
	ErrorCodeAuthorization: "Unauthorization!",
}
