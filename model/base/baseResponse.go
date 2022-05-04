package base

type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
}

func NewResponse(retcode *RetCode) *BaseResponse {
	return NewDataResponse(nil, retcode)
}

func NewDataResponse(data interface{}, retcode *RetCode) *BaseResponse {
	return &BaseResponse{
		Code:    retcode.Code,
		Message: retcode.Message,
		Data:    data,
	}
}

func NewErrorResponse(err error, retcode *RetCode) *BaseResponse {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}

	return &BaseResponse{
		Code:    retcode.Code,
		Message: retcode.Message,
		Error:   errMsg,
	}
}
