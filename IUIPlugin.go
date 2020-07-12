package interfaces

// IUI ...
type IUIPlugin interface {
	ClearScreen(i interface{}) interface{}
	PrintPercentOfScreen(i interface{}, str string, percent int) interface{}
	PrintBanner(i interface{}) interface{}
	PrintWorkloadOverview(i interface{})
	PrintTable(i interface{}, heads []string, rows [][]string) interface{}
}
