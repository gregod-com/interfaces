package interfaces

// IRegistry describes the registry and allows basic operations ...
type IRegistry interface {
	GetName() string
	GetImages() ([]IImage, error)
	GetImageByID(int) IImage
	GetStage() string
	GetID() int32
}
