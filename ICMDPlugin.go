package interfaces

// ICMDPlugin ...
type ICMDPlugin interface {
	GetCommands(c *cli.Context) []*cli.Command
}
