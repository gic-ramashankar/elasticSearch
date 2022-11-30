package helper

import (
	"encoding/json"
	"net/http"
)

const StatusCodeOK = "OK"
const StatusCodeNotOK = "NotOK"

type Response struct {
	Status string      `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data,omitempty"`
}

func (r *Response) ToJson() []byte {
	json_data, _ := json.Marshal(r)
	return json_data
}

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondWithJson(w, code, msg, StatusCodeNotOK, map[string]string{})
}

func RespondWithJson(w http.ResponseWriter, code int, msg string, status string, payload interface{}) {

	resp := Response{Data: payload, Status: status, Msg: msg}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(resp.ToJson())
}
