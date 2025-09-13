package common

type ErrorDetail struct {
	Message     string `json:"message"`
	Description string `json:"description,omitempty"`
}

type CalcioError struct {
	Code    int           `json:"code"`
	Details []ErrorDetail `json:"details,omitempty"`
}

func NewCalcioError(code int, err error) CalcioError {
	return CalcioError{
		Code: code,
		Details: []ErrorDetail{
			{Message: err.Error()},
		},
	}
}

func AppendError(calcioErr CalcioError, err error) CalcioError {
	calcioErr.Details = append(calcioErr.Details, ErrorDetail{Message: err.Error()})
	return calcioErr
}

func (e CalcioError) Error() string {
	if len(e.Details) > 0 {
		return e.Details[0].Message
	}
	return "unknown error"
}

func (e CalcioError) IsEmpty() bool {
	return len(e.Details) == 0
}

func (e *CalcioError) SetCode(code int) {
	e.Code = code
}
