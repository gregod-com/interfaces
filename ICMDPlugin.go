package interfaces

import (
	"github.com/urfave/cli/v2"
)

// ICMDPlugin ...
type ICMDPlugin interface {
	GetCommands(i interface{}) []*cli.Command
}
