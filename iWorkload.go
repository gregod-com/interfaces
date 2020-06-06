package interfaces

// IWorkload is the more concrete definition of a ...
type IWorkload interface {
	// Init(IWorkloadMetadata) error
	GetName() string
	PathToPodDescriptor() string
	GetPod() IPod
	GetActive() (string, string, bool)
}
