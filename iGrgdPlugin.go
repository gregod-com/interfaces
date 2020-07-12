package interfaces

// IGrgdPlugin ...
type IGrgdPlugin interface {
	Init(i interface{}) interface{}
	GetMetaData(i interface{}) interface{}
	Methods(i interface{}) map[string]interface{}
	GetUIPlugin(i interface{}) IUIPlugin // this might not be a good idea
}

type PluginMetadata struct {
	name     string
	version  string
	size     uint64
	url      string
	category string
}

func (m pluginMetadata) Name() string {
	return m.name
}

func (m pluginMetadata) Version() string {
	return m.version
}

func (m pluginMetadata) Size() uint64 {
	return m.size
}

func (m pluginMetadata) URL() string {
	return m.url
}

func (m pluginMetadata) Category() string {
	return m.category
}
