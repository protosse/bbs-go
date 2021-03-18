package common

type JsonResult struct {
	ErrorCode int         `json:"errorCode"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Success   bool        `json:"success"`
}

func Json(code int, message string, data interface{}, success bool) *JsonResult {
	return &JsonResult{
		ErrorCode: code,
		Message:   message,
		Data:      data,
		Success:   success,
	}
}

func JsonData(data interface{}) *JsonResult {
	return &JsonResult{
		ErrorCode: 0,
		Data:      data,
		Success:   true,
	}
}

func JsonSuccess() *JsonResult {
	return &JsonResult{
		ErrorCode: 0,
		Success:   true,
	}
}

func JsonTipError(tip string) *JsonResult {
	return &JsonResult{
		ErrorCode: TipErrorCode,
		Message:   tip,
		Success:   false,
	}
}

func JsonError(err *CodeError) *JsonResult {
	return &JsonResult{
		ErrorCode: err.Code,
		Message:   err.Message,
		Success:   false,
	}
}
