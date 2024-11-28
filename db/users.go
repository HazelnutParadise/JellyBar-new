package db

import (
	"jellybar/obj"

	"gorm.io/gorm"
)

func GetUsers() ([]obj.User, error) {
	var users []obj.User
	err := database.Find(&users).Error
	return users, err
}

func AddUser(user *obj.User) error {
	err := database.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *obj.User) error {
	err := database.Transaction(func(tx *gorm.DB) error {
		var err error
		var oldUser obj.User
		tx.Preload("Author").First(&oldUser, user.ID)
		if user.Role > obj.UserRoleUser {
			if oldUser.Author.ID == 0 {
				tx.Create(&obj.Author{UserID: user.ID})
			}
		}
		err = tx.Save(user).Error
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func DeleteUser(user *obj.User) error {
	result := database.Delete(user)
	return result.Error
}

// TODO: 跟會員系統資料庫同步，可做在該帳號每次登入時同步
