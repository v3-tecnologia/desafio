package gateway

type FaceRecognizerInterface interface {
	RecognizeFace(imagePath string) (bool, error)
}
