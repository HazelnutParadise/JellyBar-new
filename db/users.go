package db

import (
	"jellybar/obj"
)

func GetUsers() ([]obj.User, error) {
	var users []obj.User
	result := database.Find(&users)
	return users, result.Error
}

func AddUser(user obj.User) error {
	result := database.Create(&user)
	return result.Error
}

func UpdateUser(user obj.User) {
	database.Save(&user)
}

func DeleteUser(user obj.User) {
	database.Delete(&user)
}

// TODO: 跟會員系統資料庫同步，可做在該帳號每次登入時同步
