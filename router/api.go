package router

import (
	"encoding/json"
	"fmt"
	"jellybar/db"
	"jellybar/obj"
	"jellybar/utils"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
)

func defineApi(r *gin.RouterGroup) {
	r.POST("/user", func(ctx *gin.Context) {
		var user = obj.User{
			Username:       "",
			Uuid:           "",
			Name:           "",
			Role:           "USER",
			CreateAt:       time.Now(),
			Status:         "active",
			StatusUpdateAt: time.Now(),
		}
		ctx.BindJSON(&user)
		condition := map[string]string{
			"username": user.Username,
		}
		conditionJson, _ := json.Marshal(condition)
		resp, _ := http.Get("http://192.168.1.109:5004/DB/record?relation=users&conditions=" + url.QueryEscape(string(conditionJson)))
		var resultMap map[string]interface{}
		if resp.StatusCode == 200 {
			json.NewDecoder(resp.Body).Decode(&resultMap)

			if len(resultMap) == 0 {
				ctx.JSON(409, gin.H{"message": "User not exist in the system"})
				return
			}
			// 將 resultMap["result"] 轉換為 []interface{}
			results, ok := resultMap["result"].([]interface{})
			if !ok || len(results) == 0 {
				ctx.JSON(409, gin.H{"message": "User not exist in the system"})
				return
			}

			// 將第一個元素轉換為 map[string]interface{}
			resultMap2 := results[0].([]interface{})

			fmt.Printf("%s", resultMap2)

			// 確保 username 是 string 類型
			username := resultMap2[1].(string)
			if username != user.Username {
				ctx.JSON(409, gin.H{"message": "User not exist in the system"})
				return
			}

			user.Uuid = resultMap2[0].(string)
			if utils.IsASCII(resultMap2[4].(string)) {
				user.Name = resultMap2[4].(string) + " " + resultMap2[5].(string)
			} else {
				user.Name = resultMap2[5].(string) + " " + resultMap2[4].(string)
			}
		} else {
			ctx.JSON(500, gin.H{"message": "Internal server error"})
			return
		}
		db.AddUser(user)
		ctx.JSON(200, gin.H{"message": "User created"})
	})
}
