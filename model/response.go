package model

type Response struct {
	Data  interface{} `json:"data" bson:"data"`
	Error string      `json:"error" bson:"error"`
	OK    bool        `json:"ok" bson:"ok"`
}

func NewResponse(data interface{}, error string, OK bool) (response Response) {
	var r Response
	r.Data = data
	r.Error = error
	r.OK = OK
	return r
}
