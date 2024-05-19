package quickapiwendy

type (
	ErrorDTO struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
)

const (
	codeBadJson  = "JSON"
	codeBadInput = "VALIDATION"
	codeGeneric  = "GENERIC" // ;-)
)

func createError(code string, err error) *ErrorDTO {
	return &ErrorDTO{code, err.Error()}
}

func (e *ErrorDTO) Error() string {
	return e.Message
}
