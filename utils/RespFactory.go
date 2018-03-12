// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package utils

type ResponseData struct {
	Status  uint16      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var maps = map[uint16]string{
	200: "success.",
	201: "invalid param.",
	202: "illegal signature value.",
	500: "system error.",
	501: "parse obj to json fail",
}

func NewResponseData() *ResponseData {
	var data = new(ResponseData)
	data.Status = 200
	data.Message = maps[200]
	return data
}

func NewResponseDataWithStatus(status uint16) *ResponseData {
	var data = new(ResponseData)
	data.Status = status
	data.Message = maps[status]
	return data
}

func NewResponseDataWithStatusMessage(status uint16, message string) *ResponseData {
	var data = new(ResponseData)
	data.Status = status
	data.Message = message
	return data
}

func Build() *ResponseData {
	defaultResp := NewResponseData()
	return defaultResp
}

func BuildData(data interface{}) *ResponseData {
	defaultResp := NewResponseData()
	defaultResp.Data = data
	return defaultResp
}

func BuildStatus(status uint16) *ResponseData {
	defaultResp := NewResponseDataWithStatus(status)
	return defaultResp
}

func BuildAll(status uint16, message string, data interface{}) *ResponseData {
	defaultResp := NewResponseDataWithStatusMessage(status, message)
	defaultResp.Data = data
	return defaultResp
}
