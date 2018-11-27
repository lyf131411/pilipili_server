package defs

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErroResponse struct {
	HttpSc int
	Error  Err
}

var (
	ErrorRequestBodyParseFailed = ErroResponse{HttpSc: 400, Error: Err{Error: "Request body is not correct", ErrorCode: "001"}}
	ErrorNotAuthUser            = ErroResponse{HttpSc: 401, Error: Err{Error: "User authentication failed", ErrorCode: "002"}}
	ErrorDBError                = ErroResponse{HttpSc: 500, Error: Err{Error: "DB ops dailed", ErrorCode: "003"}}
	ErrorInternalFaults         = ErroResponse{HttpSc: 500, Error: Err{Error: "Internal service error", ErrorCode: "004"}}
)
