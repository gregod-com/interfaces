package interfaces

// ISortable ...
type ISortable interface {
	GetName() string
	GetValues(i ...interface{}) []string
}
