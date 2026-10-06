package handler

import (
	"main/config/jwt"
	"main/model"
	"main/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateArticle(c *gin.Context) {
	var article model.Article

	err := c.ShouldBindJSON(&article)
	if err != nil {
		c.Error(err)
		return
	}
	claims, err := jwt.GetClaims(c)
	if err != nil {
		c.Error(err)
		return
	}

	article.UserId = claims.ID

	e := service.CreateArticle(article.UserId, article)
	if e != nil {
		c.Error(e)
		return
	}
}

func GetArticleUser(c *gin.Context) {

	claims, err := jwt.GetClaims(c)
	if err != nil {
		c.Error(err)
		return
	}

	article, er := service.GetArticleUser(claims.ID)
	if er != nil {
		c.Error(er)
		return
	}
	c.JSON(200, article)
}

func GetArticle(c *gin.Context) {
	a, er := service.GetArticle()
	if er != nil {
		c.Error(er)
		return
	}
	c.JSON(200, a)

}

func GetArticleById(c *gin.Context) {
	id := c.Param("id")
	idd, err := strconv.Atoi(id)
	if err != nil {
		c.Error(err)
		return
	}
	a, er := service.GetArticleById(idd)
	if er != nil {
		c.Error(er)
		return
	}
	c.JSON(200, a)

}
func UpdatedArticle(c *gin.Context) {

	var a model.Article
	if err := c.ShouldBindJSON(&a); err != nil {
		c.Error(err)
		return
	}
	id := c.Param("id")
	idd, err := strconv.Atoi(id)
	if err != nil {
		c.Error(err)
		return
	}
	claims, err := jwt.GetClaims(c)
	if err != nil {
		c.Error(err)
		return
	}
	if err := service.UpdatedArticle(idd, claims.ID, a); err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, gin.H{
		"message": "update berhasil",
	})

}
func DeleteArticle(c *gin.Context) {
	id := c.Param("id")
	idd, err := strconv.Atoi(id)
	if err != nil {
		return
	}
	claims, err := jwt.GetClaims(c)
	if err != nil {
		c.Error(err)
		return
	}

	if err := service.DeleteArticle(idd, claims.ID); err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, gin.H{
		"message": "berhasil",
	})
}
