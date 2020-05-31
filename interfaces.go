package interfaces

import (
	"time"
)

// // IServiceWrapper ...
// type IServiceWrapper interface {
// 	// GetName() string
// 	Init(IPod) error
// 	// GetSettings() IServiceSetting
// 	GetActive() (string, string, bool)
// 	// GetPod() DockerComposePod
// 	// GetMaster() DockerComposeContainer
// 	// GetSidecars() map[string]DockerComposeContainer
// 	// GetAllContainers() map[string]DockerComposeContainer
// }

// IConfigObject interface ...
type IConfigObject interface {
	Update() error
	PrintConfig() error
	GetSourceAsString() string
	GetSourceAsBytes() []byte
	GetConfigPath() string
	IsDebug() bool
	GetProjectDir() string
	GetWorkloadMetadata() map[string]IWorkloadMetadata
	GetWorkloads() map[string]IWorkload
	GetRegistries() map[string]string
	AddWorkloadShortcut(string, string) error
	RemoveWorkloadShortcut(string) error
	GetWorkloadShortcuts() map[string]string
	GetWorkloadByShortcut(string) string
	WasCommandUsed(string) bool
	LearnedCommands() int
	MarkCommandLerned(string) error
	GetLastUsed() time.Time
}

// IConfigObject interface ...
type IPluginIndex interface {
	Update() error
	PrintConfig() error
	// GetSourceAsString() string
	GetSourceAsBytes() []byte
	GetConfigPath() string
	GetPluginList() interface{}
	// IsDebug() bool
	// GetProjectDir() string
	GetLastChecked() time.Time
}

// // IamServiceSetting ...
// type IServiceSetting interface {
// 	GetName() string
// 	SetName(string) error
// 	GetActive() bool
// 	SetActive(bool) error
// 	GetEnv() string
// 	SetEnv(string) error
// }

// IWorkloadMetadata ...
type IWorkloadMetadata interface {
	GetName() string
	GetActive() (string, string, bool)
	ToggleActive() error
	GetEnvAsString() string
	GetEnvAsEmoji() string
	GetPathToPodDescriptor() string
}

// IWorkload is the more concrete definition of a ...
type IWorkload interface {
	// Init(IWorkloadMetadata) error
	GetName() string
	PathToPodDescriptor() string
	GetPod() IPod
	GetActive() (string, string, bool)
}

// IPod ...
type IPod interface {
	GetName() string
	GetMainContainer() IContainer
	GetSidecars() []IContainer
	GetAllContainers() map[string]IContainer
}

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

// IRegistry describes the registry and allows basic operations ...
type IRegistry interface {
	GetName() string
	GetImages() ([]IImage, error)
	GetImageByID(int) IImage
	GetStage() string
	GetID() int32
}

// IImage ...
type IImage interface {
	GetFullName() string
	GetRepositoryAsString() string
	SetRepository(string) error
	GetTagAsString() string
	SetTag(string) error
	GetSize() int
	GetCreated() time.Time
}

type IRepository interface {
	GetName() string
	GetFullUrl() string
	GetAllTags() []string
	GetCategory() string
	GetDescription() string
}

type IVolume interface {
	GetHostPath() string
	GetContainerPath() string
	IsReadOnly() bool
	Delete() error
	Detach() error
	GetOwner() error
}
