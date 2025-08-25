package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func M3U(ctx *gin.Context) {
	redactedHeaders := redactSensitiveHeaders(ctx.Request.Header)
	js, err := json.MarshalIndent(redactedHeaders, "", "    ")
	if err != nil {
		fmt.Errorf("json marshal has error :%s", err)
	}
	fmt.Println(string(js))
}

// redactSensitiveHeaders returns a copy of the given header map with sensitive header values redacted.
func redactSensitiveHeaders(headers map[string][]string) map[string][]string {
	sensitiveHeaders := []string{
		"Authorization",
		"Proxy-Authorization",
		"Cookie",
		"Set-Cookie",
	}
	redacted := make(map[string][]string, len(headers))
	for k, v := range headers {
		redactedKey := false
		for _, sh := range sensitiveHeaders {
			if strings.EqualFold(k, sh) {
				redactedKey = true
				break
			}
		}
		if redactedKey {
			redacted[k] = []string{"REDACTED"}
		} else {
			redacted[k] = v
		}
	}
	return redacted
}
