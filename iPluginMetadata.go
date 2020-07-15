package interfaces

// IPluginMetadata ...
type IPluginMetadata interface {
	GetName() string
	GetVersion() string
	GetSize() uint64
	GetURL() string
	GetCategory() string
	GetActive() bool
}
