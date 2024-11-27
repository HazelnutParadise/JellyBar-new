package app

import (
	"encoding/json"
	"jellybar/db"
	"jellybar/obj"
	"jellybar/utils"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
)

type userAPIResponse struct {
	Result []struct {
		Username  string `json:"username"`
		Uuid      string `json:"uuid"`
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
	} `json:"result"`
}

func HandleGetUserList(ctx *gin.Context) {
	users := db.GetUsers()
	ctx.JSON(200, gin.H{"users": users})
}

func HandlePostUser(ctx *gin.Context) {
	var user = obj.User{
		Username: "",
		Uuid:     "",
		Name:     "",

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
	resp, _ := http.Get("http://192.168.1.109:5004/DB/record?relation=users&return_as_dict=true&conditions=" + url.QueryEscape(string(conditionJson)))

	// 使用新的結構體
	var apiResponse userAPIResponse
	if resp.StatusCode == 200 {
		if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
			ctx.JSON(500, gin.H{"message": "解析回應資料時發生錯誤"})
			return
		}

		if len(apiResponse.Result) == 0 {
			ctx.JSON(409, gin.H{"message": "榛果繽紛樂會員系統無此用戶"})
			return
		}

		userResult := apiResponse.Result[0]
		if userResult.Username != user.Username {
			ctx.JSON(409, gin.H{"message": "未知的資料不一致"})
			return
		}

		user.Uuid = userResult.Uuid
		if utils.IsASCII(userResult.Lastname) {
			user.Name = userResult.Firstname + " " + userResult.Lastname
		} else {
			user.Name = userResult.Lastname + " " + userResult.Firstname
		}
	} else {
		ctx.JSON(502, gin.H{"message": "無法連線到榛果繽紛樂會員系統"})
		return
	}
	db.AddUser(user)
	ctx.JSON(200, gin.H{"message": "用戶新增成功"})
}
