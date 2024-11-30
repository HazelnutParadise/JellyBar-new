package app

import (
	"encoding/json"
	"errors"
	"jellybar/db"
	"jellybar/obj"
	"jellybar/utils"
	"net/http"
	"net/url"
	"time"

	"github.com/HazelnutParadise/Go-Utils/conv"
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

type userFrontend struct {
	*obj.User
	Role string `json:"role"`
}

func HandleGetUserList(ctx *gin.Context) {
	var result []userFrontend
	users, err := db.GetUsers()
	if err != nil {
		ctx.JSON(500, gin.H{"message": "取得用戶列表失敗\n" + err.Error()})
		return
	}
	for _, user := range *users {
		role := convertUserRoleToString(user.Role)
		result = append(result, userFrontend{User: &user, Role: role})
	}
	ctx.JSON(200, gin.H{"users": &result})
}

func HandlePostUser(ctx *gin.Context) {
	var user = obj.User{
		Username:       "",
		Uuid:           "",
		Name:           "",
		Role:           obj.UserRoleUser,
		CreateAt:       time.Now(),
		Status:         "active",
		StatusReason:   "-",
		StatusUpdateAt: time.Now(),
	}
	ctx.BindJSON(&user)
	condition := map[string]string{
		"username": user.Username,
	}
	statusCode, err := getUserFromHazelnutParadiseDB(&user, condition)
	if err != nil {
		ctx.JSON(statusCode, gin.H{"message": err.Error()})
		return
	}
	err = db.AddUser(&user)
	if err != nil {
		ctx.JSON(500, gin.H{"message": "用戶新增失敗\n" + err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "用戶新增成功"})
}

func HandleUpdateUser(ctx *gin.Context) {
	var user obj.User
	user.ID = uint(conv.ParseInt(ctx.Param("id")))
	role := ctx.Query("role")

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"message": "解析用戶資料失敗：" + err.Error()})
		return
	}

	if role != "" {
		userRole, err := convertUserRole(role)
		if err != nil {
			ctx.JSON(400, gin.H{"message": err.Error()})
			return
		}
		user.Role = userRole
	}

	err := db.UpdateUser(&user)
	if err != nil {
		ctx.JSON(500, gin.H{"message": "用戶更新失敗\n" + err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "用戶更新成功"})
}

func HandleDeleteUser(ctx *gin.Context) {
	user := obj.User{
		ID: uint(conv.ParseInt(ctx.Param("id"))),
	}
	err := db.DeleteUser(&user)
	if err != nil {
		ctx.JSON(500, gin.H{"message": "用戶刪除失敗\n" + err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "用戶刪除成功"})
}

func getUserFromHazelnutParadiseDB(user *obj.User, condition map[string]string) (int, error) {
	conditionJson, _ := json.Marshal(condition)
	resp, _ := http.Get("http://192.168.1.109:5004/DB/record?relation=users&return_as_dict=true&conditions=" + url.QueryEscape(string(conditionJson)))
	// 使用新的結構體
	var apiResponse userAPIResponse
	if resp.StatusCode == 200 {
		if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
			return 500, errors.New("解析回應資料時發生錯誤")
		}
		if len(apiResponse.Result) == 0 {
			return 409, errors.New("榛果繽紛樂會員系統無此用戶")
		}
		userResult := &apiResponse.Result[0]

		var name string
		if utils.IsASCII(userResult.Lastname) {
			name = userResult.Firstname + " " + userResult.Lastname
		} else {
			name = userResult.Lastname + " " + userResult.Firstname
		}
		user.Uuid = userResult.Uuid
		user.Username = userResult.Username
		user.Name = name
		return 200, nil
	}
	return 502, errors.New("無法連線到榛果繽紛樂會員系統")
}

func convertUserRole(role string) (obj.UserRole, error) {
	var result obj.UserRole
	switch role {
	case "USER":
		result = obj.UserRoleUser
	case "AUTHOR":
		result = obj.UserRoleAuthor
	case "EDITOR":
		result = obj.UserRoleEditor
	case "ADMIN":
		result = obj.UserRoleAdmin
	default:
		return 255, errors.New("無效的用戶角色")
	}
	return result, nil
}

func convertUserRoleToString(role obj.UserRole) string {
	switch role {
	case obj.UserRoleUser:
		return "USER"
	case obj.UserRoleAuthor:
		return "AUTHOR"
	case obj.UserRoleEditor:
		return "EDITOR"
	case obj.UserRoleAdmin:
		return "ADMIN"
	default:
		return ""
	}
}
