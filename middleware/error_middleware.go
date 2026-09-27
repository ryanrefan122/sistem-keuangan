package middleware

import (
	"errors"
	"main/helper"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		detectederror := c.Errors
		if len(detectederror) == 0 {
			return
		}
		err := detectederror.Last()
		statusCode := http.StatusInternalServerError
		message := "Internal server error"
		var details interface{}
		// Validation error
		var validationErrs validator.ValidationErrors
		if errors.As(err, &validationErrs) {
			statusCode = http.StatusBadRequest
			message = "Validation failed"

			errorFields := make(map[string]string)

			for _, vErr := range validationErrs {
				errorFields[vErr.Field()] = getCustomMessage(vErr)
			}

			details = errorFields
		} else {
			// Application error
			var appErr *helper.AppError

			if errors.As(err, &appErr) {
				statusCode = appErr.Status
				message = appErr.Message
			}
		}

		helper.SendError(c, statusCode, message, details)
	}
}

// Helper lokal untuk menerjemahkan tag validator menjadi kalimat yang enak dibaca
func getCustomMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "Field ini wajib diisi"
	case "email":
		return "Format email tidak valid"
	case "min":
		return "Panjang minimal adalah " + fe.Param() + " karakter"
	case "max":
		return "Panjang maksimal adalah " + fe.Param() + " karakter"
	default:
		return "Field ini tidak valid"
	}
}
