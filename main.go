// Package happy_translate @Description  TODO
// @Author  	 jiangyang
// @Created  	 2023/10/13 18:14
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"happy_translate/baidutranslate"
	"log"
	"net/http"
)

func main() {
	e := gin.Default()
	g := e.Group("").Use(Logger(), Auth())
	g.GET("/t", Translate)
	if err := e.Run(":8080"); err != nil {
		log.Fatal("Exit err:", err)
	}
}

func Logger() func(*gin.Context) {
	return func(ctx *gin.Context) {
		log.Println(ctx.RemoteIP(), ctx.Request.RequestURI, ctx.Request.Body)
		ctx.Next()
	}
}

func Auth() func(*gin.Context) {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("TOKEN")
		if token != "JIANGYANG" && token != "TEST" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error_code": 10004,
				"error_msg":  "未授权访问",
			})
			ctx.Abort()
		}
		ctx.Next()
	}
}

func Translate(ctx *gin.Context) {
	q := ctx.Query("q")
	if len(q) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"error_code": 10001,
			"error_msg":  "非法空值",
		})
		return
	}
	result, err := baidutranslate.Translate(q)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusOK, gin.H{
			"error_code": 10002,
			"error_msg":  "查询失败",
		})
		return
	}
	if len(result.TransResult) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"error_code": 10003,
			"error_msg":  "结果为空",
		})
		return
	}
	log.Println(result.TransResult)
	_, _ = ctx.Writer.Write([]byte(fmt.Sprintf(`%s`, result.TransResult[0].Dst)))
}
