package util

import "encoding/json"

type ResponseJSON struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Model interface{} `json:"model"`
}

//GetResponseJSONInString ...
func GetResponseJSONInString(responseData interface{}) string {
	bytes, _ := json.Marshal(responseData)
	return string(bytes)
}
