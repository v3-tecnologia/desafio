package web

import (
	"encoding/json"
	"fmt"
	dto2 "github.com/HaroldoFV/desafio/internal/dto"
	"github.com/HaroldoFV/desafio/internal/usecase"
	"net/http"
)

type GPSHandler struct {
	CreateGPSUseCase usecase.CreateGPSUseCaseInterface
}

func NewGPSHandler(
	createGPSUsecase usecase.CreateGPSUseCaseInterface,
) *GPSHandler {
	return &GPSHandler{
		CreateGPSUseCase: createGPSUsecase,
	}
}

// Create gps godoc
// @Summary     Create	gps
// @Description Create	gps
// @Tags        gps
// @Accept      json
// @Produce     json
// @Param       request		body  		dto.CreateGPSInputDTO		true	"gps request"
// @Success     201
// @Failure     500			{object}	Error
// @Failure     400			{object}	Error
// @Router      /gps	[post]
func (h *GPSHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto dto2.CreateGPSInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		fmt.Println("Error decoding request body:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.CreateGPSUseCase.Execute(dto)
	if err != nil {
		validationErrors := map[string]bool{
			"invalid id":                             true,
			"latitude must be between -90 and 90":    true,
			"longitude must be between -180 and 180": true,
			"MAC address cannot be empty":            true,
			"invalid MAC address format":             true,
		}
		HandleError(w, err, validationErrors)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		fmt.Println("Error encoding response:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("GPS created successfully")
}
