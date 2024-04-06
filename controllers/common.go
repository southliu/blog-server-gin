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

func ReturnPageSuccess(ctx *gin.Context, code int, message interface{}, data interface{}, count int64) {
	json := &PageStruct{
		Code:    code,
		Message: message,
		Data:    data,
		Count:   count,
	}
	ctx.JSON(200, json)
}

func ReturnSuccess(ctx *gin.Context, code int, message interface{}, data interface{}) {
	json := &JsonStruct{
		Code:    code,
		Message: message,
		Data:    data,
	}
	ctx.JSON(200, json)
}

func ReturnError(ctx *gin.Context, code int, message interface{}) {
	json := &JsonErrorStruct{
		Code:    code,
		Message: message,
	}
	ctx.JSON(200, json)
}

func EncryptMd5(s string) string {
	ctx := md5.New()
	ctx.Write([]byte(s))
	return hex.EncodeToString(ctx.Sum(nil))
}
