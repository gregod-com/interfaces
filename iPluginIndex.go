package interfaces

import (
	"time"
)

// IPluginIndex interface ...
type IPluginIndex interface {
	Update() error
	PrintConfig() error
	GetConfigPath() string
	GetPluginList() []IPluginMetadata
	GetLastChecked() time.Time
	AddPlugin(newplug IPluginMetadata) error
}
