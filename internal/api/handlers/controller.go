package handlers

type DataModel interface {
}

type ApiController struct {
	db []DataModel
}

func NewApiController() *ApiController {
	return &ApiController{
		db: make([]DataModel, 0),
	}
}
