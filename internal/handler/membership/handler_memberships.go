package membership

import (
	"context"

	"github.com/Fairuzzzzz/perpustakaan-api/internal/middleware"
	"github.com/Fairuzzzzz/perpustakaan-api/internal/model/memberships"
	"github.com/gin-gonic/gin"
)

type membershipService interface {
	SignUp(ctx context.Context, req memberships.SignUpRequest) error
	Login(ctx context.Context, req memberships.LoginRequest) (string, error)
	DeleteUsers(ctx context.Context, req memberships.DeleteUserRequest) error
	GetAllUser(ctx context.Context, pageSize, pageIndex int) (memberships.GetAllUserResponse, error)
}

type Handler struct {
	*gin.Engine

	membershipSvc membershipService
}

func NewHandler(api *gin.Engine, membershipSvc membershipService) *Handler {
	return &Handler{
		Engine:        api,
		membershipSvc: membershipSvc,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("membership")
	route.POST("/sign-up", h.SignUp)
	route.POST("/login", h.Login)
	route.Use(middleware.AuthMiddleware())
	route.Use(middleware.AdminOnly())
	route.DELETE("/delete-user", h.DeleteUser)
	route.GET("/", h.GetAllUser)
}
