package gyroscope

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/kevenmiano/v3/internal/domain"
	"github.com/kevenmiano/v3/internal/infra"
	service "github.com/kevenmiano/v3/internal/service/gyroscope"
)

type Route struct {
	gyroscopeRepository service.IService
}

func (r Route) Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	gyroscopeDto := domain.GyroscopeDto{}

	err := json.Unmarshal([]byte(request.Body), &gyroscopeDto)

	if err != nil {
		return infra.ErrorResponse(http.StatusBadRequest, err.Error()), nil
	}

	gyroscope := domain.NewGyroscope(&gyroscopeDto)

	gyroscope.SetDeviceID(request.RequestContext.Identity.SourceIP)

	createdLocation, err := r.gyroscopeRepository.Create(gyroscope)

	if err != nil {
		return infra.ErrorResponse(http.StatusBadRequest, err.Error()), nil
	}

	return infra.Response(http.StatusOK, createdLocation), nil
}

func NewRoute(gyroscopeRepository service.IService) *Route {
	return &Route{
		gyroscopeRepository: gyroscopeRepository,
	}
}
