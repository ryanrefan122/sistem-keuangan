package repository

import (
	"database/sql"
	"errors"
	"main/config/database"
	"main/model"

	"github.com/go-sql-driver/mysql"
)

func CreateUser(user model.Registrasi) error {
	_, err := database.DB.Exec("INSERT INTO users(username, password, email, no_telepon ) VALUES(?,?,?,?)",
		user.Username,
		user.Password,
		user.Email,
		user.NoTelepon,
	)
	if err != nil {
		return err
	}

	return nil
}

func GetUserById(id int) (model.User, error) {

	var u model.User
	rows := database.DB.QueryRow(
		"SELECT id, username FROM users WHERE id=?",
		id,
	)
	err := rows.Scan(
		&u.ID, &u.Username,
	)
	if err != nil {
		return model.User{}, sql.ErrNoRows
	}
	return u, nil
}

func GetUser(user string) (model.User, error) {
	var u model.User
	rows := database.DB.QueryRow(
		"SELECT id, username, password FROM users WHERE username=? OR email=?",
		user, user,
	)
	err := rows.Scan(
		&u.ID, &u.Username, &u.Password,
	)
	if err != nil {
		return model.User{}, sql.ErrNoRows
	}
	return u, nil
}

func UpdateUser(id int, user model.User) error {
	_, err := database.DB.Exec("UPDATE users SET username = ?, email = ?, no_telepon = ? WHERE id = ?",
		user.Username,
		id,
	)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				return errors.New("username sudah digunakan")
			}
		}
		return err
	}
	return nil
}

func DeleteUser(id int) error {
	_, err := database.DB.Exec("DELETE FROM users WHERE id=?",
		id)
	if err != nil {
		return err
	}
	return nil
}
