
package main

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type RequestFailed struct {
	Status  string `json:"status"` // always is "failed"
	Type    string `json:"type"` // default is "Unknown"
	Message string `json:"message,omitempty"`
}

func RequestFailedFromError(err error)(*RequestFailed){
	return &RequestFailed{
		Message: err.Error(),
	}
}

func RequestFailedFromString(typ string, msg string)(*RequestFailed){
	return &RequestFailed{
		Type: typ,
		Message: msg,
	}
}

func (r *RequestFailed)MarshalJSON()(buf []byte, err error){
	res := make(gin.H, 3)
	res["status"] = "failed"
	typ := "Unknown"
	if len(r.Type) > 0 {
		typ = r.Type
	}
	res["type"] = typ
	if len(r.Message) > 0 {
		res["message"] = r.Message
	}
	return json.Marshal(res)
}

func (r *RequestFailed)SetType(t string)(*RequestFailed){
	r.Type = t
	return r
}
