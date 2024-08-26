package gps

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/kevenmiano/v3/internal/domain"
	"github.com/kevenmiano/v3/internal/infra"
	service "github.com/kevenmiano/v3/internal/service/gps"
)

type Route struct {
	gpsService service.IService
}

func (r Route) Handler(_ context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	gpsDto := domain.GPSDto{}

	err := json.Unmarshal([]byte(request.Body), &gpsDto)

	if err != nil {
		return infra.Response(http.StatusInternalServerError, fmt.Sprintf("Error unmarshalling request: %v", err)), nil
	}

	gps := domain.NewGPS(&gpsDto)

	gps.SetDeviceID(request.RequestContext.Identity.SourceIP)

	gpsCreated, err := r.gpsService.Create(gps)

	if err != nil {
		return infra.ErrorResponse(http.StatusBadRequest, fmt.Sprintf("Error creating gyroscope: %v", err)), nil
	}

	return infra.Response(http.StatusOK, gpsCreated), nil
}

func NewRoute(gpsService service.IService) *Route {
	return &Route{
		gpsService: gpsService,
	}
}
