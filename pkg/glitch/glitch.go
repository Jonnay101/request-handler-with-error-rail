package glitch

// this will be a custom error package where we can attach the required error code to the request error

// Glitch satisfies the error interface while allowing us to modify the status code
type Glitch struct {
	StatusCode   int
	ErrorMessage string
}

func (e *Error) Error() {
	return e.ErrorMessage
}

// NewError takes an error and adds a user defined error message
func NewError(err error, code int) Error {
	if code == 0 {
		code = 500
	}
	errorMessage := err.Error()
	if errorMessage == "" {
		errorMessage = "Unknown server error"
	}
	return Glitch{StatusCode: code, ErrorMessage: err.Error()}
}
