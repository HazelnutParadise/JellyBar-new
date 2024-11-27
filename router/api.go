package router

import (
	"encoding/json"
	"jellybar/db"
	"jellybar/obj"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func defineApi(r *gin.RouterGroup) {
	r.POST("/user", func(ctx *gin.Context) {
		var user obj.User
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

			stringMap := []map[string]string{}
			for _, v := range resultMap {
				stringMap = append(stringMap, v.(map[string]string))
			}

			if stringMap[0]["username"] != user.Username {
				ctx.JSON(409, gin.H{"message": "User not exist in the system"})
				return
			}
		}
		db.AddUser(user)
		ctx.JSON(200, gin.H{"message": "User created"})
	})
}
