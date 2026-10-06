package helper

import (
	"main/model"

	"golang.org/x/crypto/bcrypt"
)

func Hash(u string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(u),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", err
	}
	u = string(hash)
	return u, nil
}
func CompareHashPw(user model.Login, input model.Login) error {
	err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(input.Password),
	)
	if err != nil {
		return err
	}
	return nil

}
