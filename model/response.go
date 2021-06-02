package model

type Response struct {
	Code    int    `json:"code" bson:"code"`
	Message string `json:"message" bson:"message"`
}

func NewResponse(code int, message string) (response Response){
	var r Response
	r.Code = code
	r.Message = message
	return r
}
