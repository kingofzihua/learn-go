package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

func M3U(ctx *gin.Context) {
	js, err := json.MarshalIndent(ctx.Request.Header, "", "    ")
	if err != nil {
		fmt.Errorf("json marshal has error :%s", err)
	}
	fmt.Println(string(js))
}
