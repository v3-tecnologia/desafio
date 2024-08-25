package upload

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	adapter "github.com/kevenmiano/v3/internal/adapter/multipart"
	"github.com/kevenmiano/v3/internal/adapter/multipart/form"
	"github.com/kevenmiano/v3/internal/domain"
	"github.com/kevenmiano/v3/internal/infra"
	service "github.com/kevenmiano/v3/internal/service/photo"
	"github.com/kevenmiano/v3/internal/shared"
	"net/http"
)

type Route struct {
	photoService service.IService
}

func (m *Route) HandleMultipart(ctx context.Context, request shared.Request, decorator shared.Multipart[*domain.File]) (events.APIGatewayProxyResponse, error) {

	formData := form.NewFormData()

	multipart := adapter.NewMultipartAdapter(formData)

	file, err := multipart.Reader(request)

	if err != nil {
		return infra.ErrorResponse(http.StatusBadRequest, err.Error()), nil
	}

	return decorator(ctx, request, file)

}

func (m *Route) Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return m.HandleMultipart(ctx, request, func(ctx context.Context, request events.APIGatewayProxyRequest, d *domain.File) (events.APIGatewayProxyResponse, error) {

		photo := domain.NewPhoto(&domain.PhotoDto{
			FileName:    d.GetName(),
			Content:     d.GetContent(),
			ContentType: d.GetExt(),
		})

		createdPhoto, err := m.photoService.Create(photo)

		if err != nil {
			return infra.ErrorResponse(http.StatusBadRequest, err.Error()), nil
		}

		recognizedPhoto, err := m.photoService.Find(createdPhoto)

		if err != nil {
			return infra.ErrorResponse(http.StatusBadRequest, err.Error()), nil
		}

		return infra.Response(http.StatusOK, recognizedPhoto), nil

	})

}

func NewRoute(photoService service.IService) *Route {
	return &Route{
		photoService: photoService,
	}
}
