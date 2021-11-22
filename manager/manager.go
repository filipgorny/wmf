package manager

import (
	"github.com/filipgorny/wmf/window"
)

type Manager struct {
	windows []window.Window
}

func NewManager() *Manager {
	manager := &Manager{}

	manager.windows = make([]window.Window, 0)

	return manager
}

func (mgr *Manager) NewWindow() *window.Window {
	window := *window.NewWindow()

	mgr.windows = append(mgr.windows, window)

	return &window
}
