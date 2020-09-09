package model

// JSONResponse Using for custom response object
type JSONResponse struct {
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

func (jr JSONResponse) WithData(data interface{}) JSONResponse {
	jr.Data = data
	return jr
}
func (jr JSONResponse) WithMessage(Message string) JSONResponse {
	jr.Message = Message
	return jr
}
func (jr JSONResponse) Build() JSONResponse {
	return jr
}
