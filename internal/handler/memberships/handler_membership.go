package memberships

import (
	"context"
	"github.com/bachtiarashidiqy/simple-forum/internal/model/memberships"
	"github.com/gin-gonic/gin"
)

type membershipService interface {
	SingUp(ctx context.Context, req memberships.SingUpRequest) error
	Login(ctx context.Context, req memberships.LoginRequest) (string, error)
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
	route := h.Group("memberships")
	route.POST("/sing-up", h.SingUp)
	route.POST("/login", h.Login)
}
