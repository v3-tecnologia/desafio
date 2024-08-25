package domain

type Faces struct {
	ObjectKey string `json:"objectKey"`
}

type FacesDto struct {
	ObjectKey string
}

func NewFaces(dto *FacesDto) *Faces {
	return &Faces{
		ObjectKey: dto.ObjectKey,
	}
}
