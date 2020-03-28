package glitch

// Fail satisfies the error interface while allowing us to modify the status code
type Fail struct {
	StatusCode   int
	ErrorMessage string
}

func (e *Fail) Error() string {
	return e.ErrorMessage
}

// NewError takes an error and adds a user defined error message
func NewError(err error, code int) error {
	if code == 0 {
		code = 500
	}
	errorMessage := err.Error()
	if errorMessage == "" {
		errorMessage = "Unknown server error"
	}
	return &Fail{StatusCode: code, ErrorMessage: err.Error()}
}
