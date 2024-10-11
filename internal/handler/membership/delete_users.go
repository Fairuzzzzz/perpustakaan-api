package membership

import (
	"net/http"

	"github.com/Fairuzzzzz/perpustakaan-api/internal/model/memberships"
	"github.com/gin-gonic/gin"
)

func (h *Handler) DeleteUser(c *gin.Context) {
	ctx := c.Request.Context()

	var request memberships.DeleteUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.membershipSvc.DeleteUsers(ctx, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Status(http.StatusOK)
}
