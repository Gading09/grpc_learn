package model

type Response struct {
	Code            int         `json:"code,omitempty"`
	ErrorCode       int         `json:"errorCode,omitempty"`
	ResponseCode    string      `json:"responseCode"`
	ResponseDesc    string      `json:"responseDesc,omitempty"`
	ResponseMessage string      `json:"responseMessage,omitempty"`
	ResponseData    interface{} `json:"responseData"`
}
