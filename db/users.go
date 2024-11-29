package db

import (
	"fmt"
	"jellybar/obj"
)

func GetUsers() (*[]obj.User, error) {
	var users []obj.User
	err := database.Find(&users).Error
	return &users, err
}

func AddUser(user *obj.User) error {
	err := database.Create(user).Error
	return err
}

func UpdateUser(user *obj.User) error {
	var err error
	var oldUserAuthor obj.Author
	count := database.Where("user_id = ?", user.ID).Find(&oldUserAuthor).RowsAffected
	fmt.Println(count)
	if count == 0 {
		if user.Role > obj.UserRoleUser {
			if user.Role > obj.UserRoleUser {
				if oldUserAuthor.ID == 0 {
					user.Author.UserID = user.ID
					user.Author.Name = user.Name
				}
			}
		}
	}

	err = database.Model(user).Where("id = ?", user.ID).Updates(user).Error
	return err
}

func DeleteUser(user *obj.User) error {
	result := database.Delete(user)
	return result.Error
}

// TODO: 跟會員系統資料庫同步，可做在該帳號每次登入時同步
