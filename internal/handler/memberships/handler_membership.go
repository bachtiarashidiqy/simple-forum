package memberships

import (
	"context"
	"github.com/bachtiarashidiqy/simple-forum/internal/model/memberships"
	"github.com/gin-gonic/gin"
)

type membershipService interface {
	SingUp(ctx context.Context, req memberships.SingUpRequest) error
}

type Handler struct {
	*gin.Engine
}

func NewHandler(api *gin.Engine) *Handler {
	return &Handler{
		Engine: api,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("/memberships")
	route.GET("/sing-up", h.SingUp)
}
