package db

import (
	"jellybar/obj"

	"gorm.io/gorm"
)

func GetUsers() ([]obj.User, error) {
	var users []obj.User
	result := database.Find(&users)
	return users, result.Error
}

func AddUser(user *obj.User) error {
	err := database.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return err
		}
		user.Author.UserID = user.ID
		user.Author.Name = user.Name
		if err := tx.Create(&user.Author).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func UpdateUser(user obj.User) error {
	result := database.Save(&user)
	return result.Error
}

func DeleteUser(user obj.User) error {
	result := database.Delete(&user)
	return result.Error
}

// TODO: 跟會員系統資料庫同步，可做在該帳號每次登入時同步
