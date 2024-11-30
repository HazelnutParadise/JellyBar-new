package utils

import (
	"net/http"
	"unicode"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

// var json = jsoniter.ConfigCompatibleWithStandardLibrary
var json = jsoniter.ConfigFastest

// 自定義 JSON 函式
func FastJSON(c *gin.Context, code int, obj interface{}) {
	data, err := json.Marshal(obj)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "JSON serialization error"})
		return
	}
	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	c.Writer.WriteHeader(code)
	c.Writer.Write(data)
}

func IsASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}
