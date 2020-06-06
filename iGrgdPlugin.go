package interfaces

// IGrgdPlugin ...
type IGrgdPlugin interface {
	Init(i interface{}) interface{}
	GetMetaData(i interface{}) interface{}
	Methods(i interface{}) map[string]interface{}
}
