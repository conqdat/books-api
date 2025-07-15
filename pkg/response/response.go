package response

import "github.com/gin-gonic/gin"

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   *ErrorDetail `json:"error,omitempty"`
}

type ErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func Success(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func Error(c *gin.Context, statusCode int, message string, err error) {
	errorDetail := &ErrorDetail{
		Code:    "ERROR",
		Message: message,
	}

	if err != nil {
		errorDetail.Message = err.Error()
	}

	c.JSON(statusCode, APIResponse{
		Success: false,
		Message: message,
		Error:   errorDetail,
	})
}