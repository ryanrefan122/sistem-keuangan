package helper

import "github.com/gin-gonic/gin"

type APIResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Errors  any    `json:"errors,omitempty"` // Tempat menampung list error validasi
}

func SendError(c *gin.Context, statusCode int, message string, details interface{}) {
	c.JSON(statusCode, APIResponse{
		Success: false,
		Message: message,
		Errors:  details,
	})
}

type AppError struct {
	Status  int
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

var (
	ErrArticleNotFound = &AppError{
		Status:  404,
		Message: "Article not found",
	}

	ErrArticleForbidden = &AppError{
		Status:  403,
		Message: "You are not the owner of this article",
	}

	ErrArticleDeleted = &AppError{
		Status:  410,
		Message: "Article has been deleted",
	}
)
var (
    ErrUnauthorized = &AppError{
        Status:  401,
        Message: "Username atau password salah",
    }

    ErrUsernameExists = &AppError{
        Status:  409,
        Message: "Username sudah digunakan",
    }
)
