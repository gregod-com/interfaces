package interfaces

// IContainer ...
type IContainer interface {
	GetName() string
	SetName(string) error
	IsSideCar() bool
	IsSideCarOf() IContainer
	DefineSidecar(IContainer) error
	GetImage() IImage
	SetImage(IImage) error
	GetPorts() map[int]int
	GetPortsAsString() []string
	GetVolumes() []IVolume
	GetEnvVars() map[string]string
	AddEnvVar(string, string) error
	RemoveEnvVar(string) error
}
