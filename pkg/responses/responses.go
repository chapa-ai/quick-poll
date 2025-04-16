package responses

const (
	StatusSuccess = "success"
	StatusError   = "error"
)

type SuccessResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Error   string `json:"error"`
	Details string `json:"details,omitempty"`
}

func NewSuccessResponse(data interface{}) SuccessResponse {
	return SuccessResponse{
		Status: StatusSuccess,
		Data:   data,
	}
}

func NewErrorResponse(err string, details ...string) ErrorResponse {
	resp := ErrorResponse{
		Status: StatusError,
		Error:  err,
	}
	if len(details) > 0 {
		resp.Details = details[0]
	}
	return resp
}
