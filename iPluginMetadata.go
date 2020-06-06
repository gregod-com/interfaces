package interfaces

// IPluginMetadata ...
type IPluginMetadata interface {
	Name() string
	Version() string
	Size() uint64
	URL() string
}
