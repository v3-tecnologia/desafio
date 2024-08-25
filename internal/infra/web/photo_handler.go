package web

import (
	"encoding/json"
	"fmt"
	"github.com/HaroldoFV/desafio/internal/dto"
	"github.com/HaroldoFV/desafio/internal/usecase"
	"net/http"
)

type PhotoHandler struct {
	CreatePhotoUseCase *usecase.CreatePhotoUseCase
}

func NewPhotoHandler(createPhotoUseCase *usecase.CreatePhotoUseCase) *PhotoHandler {
	return &PhotoHandler{CreatePhotoUseCase: createPhotoUseCase}
}

var validationErrors = map[string]bool{
	"invalid id":                              true,
	"file path is required":                   true,
	"longitude must be between -180 and 180":  true,
	"MAC address cannot be empty":             true,
	"invalid MAC address format":              true,
	"invalid image type":                      true,
	"file is too large. Maximum size is 10MB": true,
}

// Create photo godoc
// @Summary     Create photo
// @Description Create photo
// @Tags        photos
// @Accept      multipart/form-data
// @Produce     json
// @Param       image formData file true "Photo file"
// @Param       mac_address formData string true "MAC Address"
// @Success     201 {object} dto.PhotoOutputDTO
// @Failure     500 {object} Error
// @Failure     400 {object} Error
// @Router      /photo [post]
func (h *PhotoHandler) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB max
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	file, fileHeader, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	macAddress := r.FormValue("mac_address")
	if macAddress == "" {
		http.Error(w, "mac_address is required", http.StatusBadRequest)
		return
	}

	input := dto.CreatePhotoInputDTO{
		Image:      fileHeader,
		MacAddress: macAddress,
		FileSize:   fileHeader.Size,
	}

	output, err := h.CreatePhotoUseCase.Execute(input)
	if err != nil {
		HandleError(w, err, validationErrors)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Photo created successfully")
}
