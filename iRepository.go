package interfaces

// IRepository ...
type IRepository interface {
	GetName() string
	GetFullUrl() string
	GetAllTags() []string
	GetCategory() string
	GetDescription() string
}
