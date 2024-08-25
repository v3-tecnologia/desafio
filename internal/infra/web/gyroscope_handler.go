package web

import (
	"encoding/json"
	"fmt"
	"github.com/HaroldoFV/desafio/internal/domain"
	dto2 "github.com/HaroldoFV/desafio/internal/dto"
	"github.com/HaroldoFV/desafio/internal/usecase"
	"net/http"
)

type GyroscopeHandler struct {
	CreateGyroscopeUseCase *usecase.CreateGyroscopeUseCase
	GyroscopeRepository    domain.GyroscopeRepositoryInterface
}

func NewGyroscopeHandler(
	createGyroscopeUsecase *usecase.CreateGyroscopeUseCase,
	gyroscopeRepository domain.GyroscopeRepositoryInterface,
) *GyroscopeHandler {
	return &GyroscopeHandler{
		CreateGyroscopeUseCase: createGyroscopeUsecase,
		GyroscopeRepository:    gyroscopeRepository,
	}
}

// Create gyroscope godoc
// @Summary     Create	gyroscope
// @Description Create	gyroscope
// @Tags        gyroscopes
// @Accept      json
// @Produce     json
// @Param       request		body  		dto.CreateGyroscopeInputDTO		true	"gyroscope request"
// @Success     201
// @Failure     500			{object}	Error
// @Failure     400			{object}	Error
// @Router      /gyroscope	[post]
func (h *GyroscopeHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto dto2.CreateGyroscopeInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		fmt.Println("Error decoding request body:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.CreateGyroscopeUseCase.Execute(dto)
	if err != nil {
		validationErrors := map[string]bool{
			"invalid id":           true,
			"name cannot be empty": true,
			"name cannot be longer than 100 characters": true,
			"model cannot be empty":                     true,
			"model cannot be longer than 50 characters": true,
			"MAC address cannot be empty":               true,
			"invalid MAC address format":                true,
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

	fmt.Println("Gyroscope created successfully")
}
