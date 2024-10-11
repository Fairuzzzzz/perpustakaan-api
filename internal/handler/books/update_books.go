package books

import (
	"net/http"

	"github.com/Fairuzzzzz/perpustakaan-api/internal/model/books"
	"github.com/gin-gonic/gin"
)

func (h *Handler) UpdateBook(c *gin.Context) {
	ctx := c.Request.Context()

	var request books.UpdateBookRequest
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.bookSvc.UpdateBook(ctx, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Status(http.StatusOK)
}
