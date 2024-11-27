package db

import "jellybar/obj"

func GetUsers() []obj.User {
	var users []obj.User
	database.Find(&users)
	return users
}

func AddUser(user obj.User) {
	database.Create(&user)
}

func UpdateUser(user obj.User) {
	database.Save(&user)
}

func DeleteUser(user obj.User) {
	database.Delete(&user)
}
