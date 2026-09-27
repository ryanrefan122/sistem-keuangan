package service

import (
	"main/model"
	"main/repository"
)

func GetProfil(id int) (model.User, error) {

	user, err := repository.GetUserById(id)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
func UpdateUser(id int, user model.User) error {
	err := repository.UpdateUser(id, user)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(id int) error {

	err := repository.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
