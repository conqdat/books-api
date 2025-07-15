package handlers

import (
	"net/http"

	"github.com/conqdat/books-api/internal/service"
	"github.com/conqdat/books-api/pkg/response"
	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	bookService *service.BookService
}

func NewBookHandler(bookService *service.BookService) *BookHandler {
	return &BookHandler{
		bookService: bookService,
	}
}

func (h *BookHandler) GetBooks(c *gin.Context) {
	books, err := h.bookService.GetAllBooks(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to get books", err)
		return
	}

	response.Success(c, http.StatusOK, "Books retrieved successfully", books)
}