package window

type Window struct {
	Id int64
	Width uint16
	Height uint16
}

var currentId int64 = 0

func NewWindow() *Window {
	currentId++
	return &Window{Id: currentId}
}
