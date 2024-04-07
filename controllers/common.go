package controllers

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/gin-gonic/gin"
)

type PageStruct struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
	Count   int64       `json:"count"`
}

type JsonStruct struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

type JsonErrorStruct struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

func ReturnPageSuccess(c *gin.Context, code int, message interface{}, data interface{}, count int64) {
	json := &PageStruct{
		Code:    code,
		Message: message,
		Data:    data,
		Count:   count,
	}
	c.JSON(200, json)
}

func ReturnSuccess(c *gin.Context, code int, message interface{}, data interface{}) {
	json := &JsonStruct{
		Code:    code,
		Message: message,
		Data:    data,
	}
	c.JSON(200, json)
}

func ReturnError(c *gin.Context, code int, message interface{}) {
	json := &JsonErrorStruct{
		Code:    code,
		Message: message,
	}
	c.JSON(200, json)
}

func EncryptMd5(s string) string {
	c := md5.New()
	c.Write([]byte(s))
	return hex.EncodeToString(c.Sum(nil))
}
