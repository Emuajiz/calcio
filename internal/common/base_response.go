package common

type BaseResponse[T any] struct {
	Message string   `json:"message"`
	Error   []string `json:"error,omitempty"`
	Data    T        `json:"data,omitempty"`
}

func ConstructResponse[T any](message string, data T, errors []string) BaseResponse[T] {
	return BaseResponse[T]{
		Message: message,
		Data:    data,
		Error:   errors,
	}
}

func ConstrutErrorResponse[T any](message string, err CalcioError) BaseResponse[T] {
	var emptyData T

	var errors []string
	for _, detail := range err.Details {
		errors = append(errors, detail.Message)
	}

	return BaseResponse[T]{
		Message: message,
		Data:    emptyData,
		Error:   errors,
	}
}

func ConstructSuccessResponse[T any](message string, data T) BaseResponse[T] {
	return BaseResponse[T]{
		Message: message,
		Data:    data,
	}
}
