package service

import (
	"database/sql"
	"errors"
	"fmt"
	"main/config/jwt"
	"main/helper"
	"main/model"
	"main/repository"
	"main/validation"

	"github.com/go-sql-driver/mysql"
)

func Login(u model.Login) error {
	err := validation.Validate.Struct(u)
	if err != nil {
		return err
	}
	user, er := repository.GetUser(u.Login)
	if er != nil {
		if errors.Is(er, sql.ErrNoRows) {
    	return helper.ErrUnauthorized
		}
		return er
	}
	st := helper.CompareHashPw(user, u)
	if st != nil {
		return helper.ErrUnauthorized
	}

	token, e := jwt.GenerateToken(user)
	if e != nil {
		return e
	}
	fmt.Println("token:", token)
	return nil
}

func Register(u model.Registrasi) error {

	if ok := validation.Validate.Struct(u); ok != nil {
		return ok
	}
	hash, st := helper.Hash(u.Password)
	if st != nil {
		return st
	}
	u.Password = hash
	err := repository.CreateUser(u)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				return errors.New(mysqlErr.Message)
			}
		}
		return err
	}
	return nil
}
