package interfaces

// IVolume ...
type IVolume interface {
	GetHostPath() string
	GetContainerPath() string
	IsReadOnly() bool
	Delete() error
	Detach() error
	GetOwner() error
}
