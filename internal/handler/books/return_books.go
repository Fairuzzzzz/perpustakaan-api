package books

import (
	"net/http"

	"github.com/Fairuzzzzz/perpustakaan-api/internal/model/books"
	"github.com/gin-gonic/gin"
)

func (h *Handler) ReturnBook(c *gin.Context) {
	ctx := c.Request.Context()

	var request books.ReturnBookRequest
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "userID not found in context",
		})
		return
	}

	userIDInt, ok := userID.(int64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "userID is not of type int64",
		})
		return
	}

	request.UserID = userIDInt

	err := h.bookSvc.ReturnBook(ctx, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Status(http.StatusOK)
}
