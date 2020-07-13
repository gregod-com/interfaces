package interfaces

// IGrgdPlugin ...
type IGrgdPlugin interface {
	Init(i interface{}) interface{}
	GetMetaData(i interface{}) interface{}
	Methods(i interface{}) map[string]interface{}
}

// PluginMetadata ...
type PluginMetadata struct {
	name     string
	version  string
	size     uint64
	url      string
	category string
}

func (m PluginMetadata) Name() string {
	return m.name
}

func (m PluginMetadata) Version() string {
	return m.version
}

func (m PluginMetadata) Size() uint64 {
	return m.size
}

func (m PluginMetadata) URL() string {
	return m.url
}

func (m PluginMetadata) Category() string {
	return m.category
}

func (m *PluginMetadata) SetName(name string) {
	m.name = name
}

func (m *PluginMetadata) SetVersion(version string) {
	m.version = version
}

func (m *PluginMetadata) SetSize(size uint64) {
	m.size = size
}

func (m *PluginMetadata) SetURL(url string) {
	m.url = url
}

func (m *PluginMetadata) SetCategory(category string) {
	m.category = category
}
