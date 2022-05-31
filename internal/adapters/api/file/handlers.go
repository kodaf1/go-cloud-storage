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
	router.GET("/files/:uuid", h.GetFileInfo)
	router.POST("/files", h.UploadFile)
}

func (h *handler) GetFileInfo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "form/json")

	fileInfo := h.fileService.GetFile(context.Background(), params.ByName("uuid"))
	if fileInfo == nil {
		logging.GetLogger().Debug("file not found")
		http.Error(w, "Not Found", 404)
		return
	}

	fileJson, err := json.Marshal(fileInfo)
	if err != nil {
		logging.GetLogger().Error("can't marshal fileInfo")
		http.Error(w, "Server Internal Error", 500)
		return
	}

	w.WriteHeader(200)
	fmt.Fprint(w, string(fileJson))
	logging.GetLogger().Debug("fileInfo returned")
}

func (h *handler) UploadFile(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, "Incorrect data")
		return
	}

	files, ok := r.MultipartForm.File["file"]
	if !ok || len(files) == 0 {
		w.WriteHeader(400)
		fmt.Fprint(w, "Incorrect data")
		return
	}

	fileInfo := files[0]
	fileReader, err := fileInfo.Open()

	dto := file.UploadFileDTO{
		Name:   fileInfo.Filename,
		Size:   fileInfo.Size,
		Reader: fileReader,
	}

	fileObject, err := h.fileService.UploadFile(context.Background(), &dto)

	fileJson, err := json.Marshal(fileObject)
	if err != nil {
		logging.GetLogger().Error("can't marshal fileInfo")
		http.Error(w, "Server Internal Error", 500)
		return
	}

	fmt.Fprint(w, string(fileJson))
	logging.GetLogger().Debug("file uploaded")
}
