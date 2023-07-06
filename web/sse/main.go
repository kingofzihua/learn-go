package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// @see https://www.ruanyifeng.com/blog/2017/05/server-sent_events.html
func main() {
	r := gin.Default()

	r.GET("/stream", func(c *gin.Context) {
		c.Header("Content-Type", "text/event-stream")
		c.Header("Cache-Control", "no-cache")
		c.Header("Connection", "keep-alive")
		c.Header("Access-Control-Allow-Origin", "*")

		ticker := time.NewTicker(5 * time.Second)

		defer ticker.Stop()

		login := false

		for {
			select {
			case <-ticker.C:

				_, _ = c.Writer.Write([]byte("retry: 10000\n"))

				data := fmt.Sprintf("data:%v\n\n", time.Now().Unix())
				_, _ = c.Writer.Write([]byte(data))

				if !login {
					_, _ = c.Writer.Write([]byte("event: login\n"))
					_, _ = c.Writer.Write([]byte("data: {\"name\":\"kingofzihua\"}\n\n"))
					login = true
				}

				// 手动刷新响应到客户端
				if f, ok := c.Writer.(http.Flusher); ok {
					f.Flush()
				}

			case <-c.Writer.CloseNotify():
				fmt.Printf("client close")
				return
			}
		}

	})

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, map[string]string{
			"test": "1",
		})
	})
	r.Run(":80")
}
