package interfaces

import (
	"time"
)

// IImage ...
type IImage interface {
	GetFullName() string
	GetRepositoryAsString() string
	SetRepository(string) error
	GetTagAsString() string
	SetTag(string) error
	GetSize() int
	GetCreated() time.Time
}
