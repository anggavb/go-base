package utils_response

type JSONResponse struct {
	Message string `json:"message"`
	Errors  any    `json:"errors,omitempty"`
	Data    any    `json:"data,omitempty"`
	Meta    any    `json:"meta,omitempty"`
}

func NewResponse(message string, data, meta, errors any) JSONResponse {
	return JSONResponse{
		Message: message,
		Errors:  errors,
		Data:    data,
		Meta:    meta,
	}
}
