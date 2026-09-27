package validation

import (
	"errors"
	"main/model"
)

func ValidationUser(u model.User) error {

	if u.Username == "" {
		return errors.New("nama tidak boleh kosong")
	}
	if u.Password == "" {
		return errors.New("password tidak boleh kosong")
	}
	if len(u.Password) < 8 {
		return errors.New("password harus lebih dari 8 karakter")
	}
	return nil
}
