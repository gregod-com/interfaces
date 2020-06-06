package interfaces

import (
	"time"
)

// IConfigObject interface ...
type IConfigObject interface {
	Update() error
	PrintConfig() error
	GetSourceAsString() string
	// GetSourceAsBytes() []byte
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
