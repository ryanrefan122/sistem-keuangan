package handler

import (
	"main/model"
	"main/service"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var u model.Login
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.Error(err)
		return
	}
	if err = service.Login(u); err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, "login berhasil")

}
func Register(c *gin.Context) {

	var u model.Registrasi
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.Error(err)
		return
	}
	if cek := service.Register(u); cek != nil {
		c.Error(cek)
		return
	}
	c.JSON(200, "registrasi berhasil")

}
