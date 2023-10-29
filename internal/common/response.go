package common

type AppStatus int

const (
	Ok     AppStatus = 0
	Failed AppStatus = 2
)

type AppResponse struct {
	Data    interface{} `json:"data"`
	Status  AppStatus   `json:"status"`
	Message string      `json:"message"`
	Error   error       `json:"error"`
}

func SendResponse(data interface{}, status AppStatus, message string, err error) *AppResponse {
	return &AppResponse{Data: data, Status: status, Message: message, Error: err}
}

func SendOkResponse(data interface{}) *AppResponse {
	return SendResponse(data, Ok, "", nil)
}

func SendFailedResponse(data interface{}, message string, error error) *AppResponse {
	return SendResponse(data, Failed, message, error)
}
