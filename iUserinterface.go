package interfaces

// IUI ...
type IUI interface {
	ClearScreen(i interface{}) interface{}
	PrintPercentOfScreen(i interface{}, str string, percent int) interface{}
	PrintBanner(i interface{}) interface{}
	PrintWorkloadOverview(i interface{})
}
