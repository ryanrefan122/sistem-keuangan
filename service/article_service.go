package service

import (
	
	"fmt"
	"main/model"
	"main/repository"
	"main/validation"
)



func GetArticle() ([]model.Article, error) {
	art, err := repository.GetArticle()
	if err != nil {
		return []model.Article{}, err
	}
	fmt.Println("artikel", art)
	return art, nil
}

func GetArticleUser(id int) ([]model.Article, error) {
	art, err := repository.GetArticleUser(id)
	if err != nil {
		return nil, err
	}

	return art, nil
}
func GetArticleById(id int) (model.Article, error) {
	art, err := repository.GetArticleById(id)
	if err != nil {
		return model.Article{}, err
	}
	return art, nil
}

func CreateArticle(id int, a model.Article) error {
	if err := validation.Validate.Struct(a); err != nil {
		return err
	}
	err := repository.CreateArticle(id, a)
	if err != nil {
		return err
	}
	return nil
}

func UpdatedArticle(id int, userid int, a model.Article) error {
	if err := validation.Validate.Struct(a); err != nil {
		return err
	}
	err := repository.UpdatedArticle(id, userid, a)
	if err != nil {
		return err
	}
	return nil
}
func DeleteArticle(id int, userid int) error {
	err := repository.DeleteArt(id, userid)
	if err != nil {
		return err
	}
	return nil
}
