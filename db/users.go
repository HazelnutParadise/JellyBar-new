package db

import (
	"jellybar/obj"
)

func GetUsers() []obj.User {
	var users []obj.User
	database.Find(&users)
	return users
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

// TODO: 跟會員系統資料庫同步
