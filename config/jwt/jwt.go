package jwt

import (
	"errors"
	"main/model"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Login struct {
	Token string `json:"token"`
}

func secretKey() []byte {
    return []byte(os.Getenv("JWT_SECRET"))
}

func GenerateToken(user model.Login) (string, error) {
	claims := jwt.MapClaims{
		"id":       user.ID,
		"username": user.Login,
	}
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)
	tokenstring, err := token.SignedString(secretKey())
	if err != nil {
		return "", err
	}
	return tokenstring, nil
}

func ParseToken(tokenStr string) (model.Claims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("metode signing tidak valid")
		}
		return secretKey(), nil
	})
	if err != nil {
		return model.Claims{}, err
	}
	if !token.Valid {
		return model.Claims{}, errors.New("token tidak valid")
	}

	jwtClaims := token.Claims.(jwt.MapClaims)

	claims := model.Claims{
		ID:       int(jwtClaims["id"].(float64)),
		Username: jwtClaims["username"].(string),
	}
	return claims, nil
}

func GetClaims(c *gin.Context) (model.Claims, error) {

	claimsValue, exists := c.Get("claims")
	claims, er := claimsValue.(model.Claims)
	if !er || !exists {
		return model.Claims{}, errors.New("model gagal di ambil")
	}
	return claims, nil
}
