package interfaces

import (
	"time"
)

// IConfigObject interface ...
type IConfigObject interface {
	Update() error
	GetSourceAsString() string
	GetConfigPath() string
	GetProjectDir() string
	WasCommandUsed(string) bool
	LearnedCommands() int
	MarkCommandLerned(string) error
	GetLastUsed() time.Time

	// those should maybe be moved to a interface that focuses on workloads
	GetWorkloadMetadata() map[string]IWorkloadMetadata
	GetWorkloads() map[string]IWorkload
	GetRegistries() map[string]string
	AddWorkloadShortcut(string, string) error
	RemoveWorkloadShortcut(string) error
	GetWorkloadShortcuts() map[string]string
	GetWorkloadByShortcut(string) string
}
