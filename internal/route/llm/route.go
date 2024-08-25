package llm

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/kevenmiano/v3/internal/domain"
	service "github.com/kevenmiano/v3/internal/service/rekognition"
	"sync"
)

type Route struct {
	rekognitionService service.IService
}

func (r Route) Handler(ctx context.Context, event events.S3Event) (events.S3Event, error) {

	var wg sync.WaitGroup

	errChan := make(chan error, len(event.Records))

	for _, record := range event.Records {
		wg.Add(1)

		go func(record events.S3EventRecord) {
			defer wg.Done()

			facesDto := &domain.FacesDto{
				ObjectKey: record.S3.Object.Key,
			}

			faces := domain.NewFaces(facesDto)

			_, err := r.rekognitionService.Create(faces)

			if err != nil {
				errChan <- err
				return
			}

		}(record)
	}

	wg.Wait()

	for i := 0; i < len(errChan); i++ {
		if err := <-errChan; err != nil {
			return events.S3Event{}, err
		}
	}

	return event, nil

}

func NewRoute(rekognitionService service.IService) Route {
	return Route{
		rekognitionService: rekognitionService,
	}
}
