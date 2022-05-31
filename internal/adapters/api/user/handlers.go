package user

import (
	"github.com/kodaf1/go-cloud-storage/internal/adapters/api"
	"github.com/kodaf1/go-cloud-storage/internal/domain/user"

	"github.com/julienschmidt/httprouter"
)

type handler struct {
	userService user.Service
}

func NewHandler(service user.Service) api.Handler {
	return &handler{userService: service}
}

func (h *handler) Register(router *httprouter.Router) {
	panic("implement me")
}
