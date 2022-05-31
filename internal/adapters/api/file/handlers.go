package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kodaf1/go-cloud-storage/internal/adapters/api"
	"github.com/kodaf1/go-cloud-storage/internal/domain/file"
	"github.com/kodaf1/go-cloud-storage/pkg/logging"

	"github.com/julienschmidt/httprouter"
)

type handler struct {
	fileService file.Service
}

func NewHandler(service file.Service) api.Handler {
	return &handler{fileService: service}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET("/file/:uuid", h.GetFile)
}

func (h *handler) GetFile(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	file, err := h.fileService.GetFile(context.Background(), params.ByName("uuid"))
	if err != nil {
		logging.GetLogger().Debug("file not found")
		http.Error(w, "Not Found", 404)
	}

	jsonFile, err := json.Marshal(file)
	if err != nil {
		logging.GetLogger().Error("can't marshal file")
		http.Error(w, "Server Internal Error", 500)
	}

	w.WriteHeader(200)
	fmt.Fprint(w, jsonFile)
	logging.GetLogger().Debug("file returned")
}

func (h *handler) UploadFile(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}
