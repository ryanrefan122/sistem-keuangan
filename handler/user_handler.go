package handler

import (
	"main/config/jwt"
	"main/model"

	"main/service"

	"github.com/gin-gonic/gin"
)

func ProfilUser(c *gin.Context) {
	claims, err := jwt.GetClaims(c)
	if err != nil {
		c.Error(err)
		return
	}

	user, err1 := service.GetProfil(claims.ID)
	if err1 != nil {
		c.Error(err1)
		return
	}
	c.JSON(200, user)
}

func UpdateUser(c *gin.Context) {
	var u model.User
	claims, err := jwt.GetClaims(c)
	if err != nil {
		c.Error(err)
		return
	}
	er := c.ShouldBindJSON(&u)
	if er != nil {
		c.Error(er)
		return
	}
	ok := service.UpdateUser(claims.ID, u)
	if ok != nil {
		c.Error(ok)
		return
	}

}

func DeleteUser(c *gin.Context) {
	claims, err := jwt.GetClaims(c)
	if err != nil {
		c.Error(err)
		return
	}
	ok := service.DeleteUser(claims.ID)
	if ok != nil {
		c.Error(ok)
		return
	}
	c.JSON(200, gin.H{
		"status":  "ok",
		"message": "berhasil",
	})
}
