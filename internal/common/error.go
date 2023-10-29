package common

type AppError struct {
	RootErr error  `json:"-"`
	Message string `json:"message"`
}

func (e *AppError) Error() string {
	return e.RootErr.Error()
}

func NewAppError(root error, msg string) *AppError {
	return &AppError{
		RootErr: root,
		Message: msg,
	}
}
