package user

import (
	"github.com/kodaf1/go-cloud-storage/internal/adapters/api"
	"github.com/kodaf1/go-cloud-storage/internal/domain/file"

	"github.com/julienschmidt/httprouter"
)

type handler struct {
	fileService file.Service
}

func NewHandler(service file.Service) api.Handler {
	return &handler{fileService: service}
}

func (h *handler) Register(router *httprouter.Router) {
	panic("implement me")
}

func (h *handler) UploadFile(router *httprouter.Router) {
	panic("implement me")
}

func (h *handler) GetFile(router *httprouter.Router) {
	panic("implement me")
}
