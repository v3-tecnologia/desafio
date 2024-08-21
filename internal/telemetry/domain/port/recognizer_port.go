package port

type DetectedFace struct {
	Confidence float64
}

type RecognizerPort interface {
	Recognize(imageURL string) ([]DetectedFace, error)
}
