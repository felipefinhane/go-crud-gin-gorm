package models

import "github.com/felipefinhane/go-crud-gin-gorm/config"

//Fetch all user data
func GetAllUsers(user *[]User) error {
	if err := config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

//Fetch only one user by ID
func GetUserByID(user *User, id string) error {
	if err := config.DB.Where("id=?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

//Insert new user data
func CreateUser(user *User) error {
	if err := config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

//Update user data
func UpdateUser(user *User, id string) error {
	if err := config.DB.Save(user).Error; err != nil {
		return err
	}
	return nil
}

//Delete user data
func DeleteUser(user *User, id string) error {
	if err := config.DB.Where("id=?", id).Delete(user).Error; err != nil {
		return err
	}
	return nil
}
