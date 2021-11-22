package xgb

import (
	"fmt"

	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/filipgorny/wmf/window"
)

const (
	borderWidth = 2
	titleBarHeight = 20
)

type XgbDriver struct {
	setup  *xproto.SetupInfo
	screen *xproto.ScreenInfo
	con    *xgb.Conn

	windows []XgbWindowId
}

type XgbWindowId struct {
	window     window.Window
	xprotoWindow xproto.Window
}

func NewGgbDriver() *XgbDriver {

	driver := XgbDriver{}

	X, err := xgb.NewConn()
	if err != nil {
		fmt.Println(err)
	}

	driver.con = X

	// xproto.Setup retrieves the Setup information from the setup bytes
	// gathered during connection.
	driver.setup = xproto.Setup(X)

	// This is the default screen with all its associated info.
	driver.screen = driver.setup.DefaultScreen(X)

	return &driver
}

func (drv *XgbDriver) PaintWindow(window window.Window) {
	wid := createWindow(drv.con, drv.screen, window.Width, window.Height);

	xgbWindow := XgbWindowId{}
	xgbWindow.window = window
	xgbWindow.xprotoWindow = wid

	drv.windows = append(drv.windows, xgbWindow)

	
}

func (drv *XgbDriver) CreateContainerWindow(xgbWindowId *XgbWindowId {
	parent := createWindow(
		drv.con,
		drv.screen,
		xgbWindowId.window.Width + borderWidth, 
		xgbWindowId.window.Height + borderWidth + titleBarHeight
	)

	xgb.ReparentWindow(drv.con, xgbWindowId.xprotoWindow, parent, 0, 0)
}

func (drv *XgbDriver) Run() {
	// Start the main event loop.
	for {
		// WaitForEvent either returns an event or an error and never both.
		// If both are nil, then something went wrong and the loop should be
		// halted.
		//
		// An error can only be seen here as a response to an unchecked
		// request.
		ev, xerr := drv.con.WaitForEvent()
		if ev == nil && xerr == nil {
			fmt.Println("Both event and error are nil. Exiting...")
			return
		}

		if ev != nil {
			fmt.Printf("Event: %s\n", ev)
		}
		if xerr != nil {
			fmt.Printf("Error: %s\n", xerr)
		}
	}
}
