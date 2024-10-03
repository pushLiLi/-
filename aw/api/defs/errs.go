package defs

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"` //001 002 003

}

type ErrorResponse struct {
	HttpStatusCode int
	Error          Err //继承了上面的err
}

var (
	//实例化两个err
	ErrorRequestBodyParseFailed = ErrorResponse{HttpStatusCode: 400, Error: Err{
		Error:     "Request body is not correct",
		ErrorCode: "string",
	}}
	ErrorNotAuthUser = ErrorResponse{HttpStatusCode: 401, Error: Err{
		//没有认证的user
		Error:     "User authenticated failed",
		ErrorCode: "002",
	}}
	ErrorDBFailed = ErrorResponse{HttpStatusCode: 500, Error: Err{
		Error:     "Database ops failed",
		ErrorCode: "003",
	}}
	ErrorInternalError = ErrorResponse{HttpStatusCode: 500, Error: Err{
		Error:     "Internal service error",
		ErrorCode: "004",
	}}
)
