package repository

import (
	"database/sql"
	"main/config/database"
	"main/model"
)

func CreateArticle(id int, a model.Article) error {
	result, err := database.DB.Exec("INSERT INTO articles(user_id, title, content) VALUES(?,?,?)",
		id, a.Judul, a.Content)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func GetArticle() ([]model.Article, error) {
	rows, err := database.DB.Query(
		"SELECT  id, user_id, title, content, created_at, updated_at FROM articles",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []model.Article

	for rows.Next() {
		var article model.Article

		err := rows.Scan(
			&article.ID,
			&article.UserId,
			&article.Judul,
			&article.Content,
			&article.CreatedAt,
			&article.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		articles = append(articles, article)
	}
	ok := rows.Err()
	if ok != nil {
		return nil, ok
	}
	return articles, nil

}

func GetArticleUser(id int) ([]model.Article, error) {
	rows, err := database.DB.Query(
		"SELECT id, user_id, title, content, created_at, updated_at FROM articles WHERE user_id =?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []model.Article

	for rows.Next() {
		var article model.Article

		err := rows.Scan(
			&article.ID,
			&article.UserId,
			&article.Judul,
			&article.Content,
			&article.CreatedAt,
			&article.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		articles = append(articles, article)
	}
	ok := rows.Err()
	if ok != nil {
		return nil, ok
	}
	return articles, nil

}
func GetArticleById(id int) (model.Article, error) {
	row := database.DB.QueryRow("SELECT id, user_id, title, content, created_at, updated_at FROM articles WHERE id = ? ", id)
	var article model.Article
	err := row.Scan(
		&article.ID,
		&article.UserId,
		&article.Judul,
		&article.Content,
		&article.CreatedAt,
		&article.UpdatedAt)
	if err != nil {
		return model.Article{}, err
	}
	return article, nil
}

func UpdatedArticle(id int, userid int, a model.Article) error {
	_, err := database.DB.Exec("UPDATE articles SET title = ?, content = ? WHERE id = ? AND user_id = ?", a.Judul, a.Content, id, userid)
	if err != nil {
		return err
	}
	return nil

}

func DeleteArt(id int, userid int) error {
	_, err := database.DB.Exec("DELETE FROM articles WHERE id = ? AND user_id = ?", id, userid)
	if err != nil {
		return err
	}
	return nil
}
