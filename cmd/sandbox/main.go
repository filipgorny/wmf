package main

import (
	"github.com/filipgorny/wmf/manager"
	"github.com/filipgorny/wmf/xgb"
)

func main() {
	driver := xgb.NewGgbDriver()

	manager := manager.NewManager()

	driver.PaintWindow(*manager.NewWindow())

	driver.Run()
}
