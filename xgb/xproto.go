package xgb

import (
	"fmt"
	"image"

	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil/xgraphics"
	"github.com/filipgorny/wmf/theme"
	"github.com/llgcode/draw2d/draw2dimg"
)

func createWindow(xcon *xgb.Conn, screen *xproto.ScreenInfo, width uint16, height uint16) *xproto.Window{
	wid, _ := xproto.NewWindowId(xcon)

	xproto.CreateWindow(xcon, screen.RootDepth, wid, screen.Root,
		0, 0, width, height, 0,
		xproto.WindowClassInputOutput, screen.RootVisual, 0, []uint32{})

	// This call to ChangeWindowAttributes could be factored out and
	// included with the above CreateWindow call, but it is left here for
	// instructive purposes. It tells X to send us events when the 'structure'
	// of the window is changed (i.e., when it is resized, mapped, unmapped,
	// etc.) and when a key press or a key release has been made when the
	// window has focus.
	// We also set the 'BackPixel' to white so that the window isn't butt ugly.
	xproto.ChangeWindowAttributes(xcon, wid,
		xproto.CwBackPixel|xproto.CwEventMask,
		[]uint32{ // values must be in the order defined by the protocol
			0xffffffff,
			xproto.EventMaskStructureNotify |
				xproto.EventMaskKeyPress |
				xproto.EventMaskKeyRelease})

	// MapWindow makes the window we've created appear on the screen.
	// We demonstrated the use of a 'checked' request here.
	// A checked request is a fancy way of saying, "do error handling
	// synchronously." Namely, if there is a problem with the MapWindow request,
	// we'll get the error *here*. If we were to do a normal unchecked
	// request (like the above CreateWindow and ChangeWindowAttributes
	// requests), then we would only see the error arrive in the main event
	// loop.
	//
	// Typically, checked requests are useful when you need to make sure they
	// succeed. Since they are synchronous, they incur a round trip cost before
	// the program can continue, but this is only going to be noticeable if
	// you're issuing tons of requests in succession.
	//
	// Note that requests without replies are by default unchecked while
	// requests *with* replies are checked by default.
	err := xproto.MapWindowChecked(xcon, wid).Check()
	if err != nil {
		fmt.Printf("Checked Error for mapping window %d: %s\n", wid, err)
	} else {
		fmt.Printf("Map window %d successful!\n", wid)
	}

	return &wid
}

func createBorder(driver XgbDriver, win xproto.Window, width uint16, height uint16) {
	r := image.Rectangle{
		Min: image.Point{0, 0},
		Max: image.Point{int(width), titleBarHeight},
	}

	dest := image.NewRGBA(r)
	gc := draw2dimg.NewGraphicContext(dest)

	t := theme.NewTheme("basic")

	source, err := draw2dimg.LoadFromPngFile(t.GetTitleBarImagePath())

	if err != nil {
		panic(err)
	}

	for i := 0; i < int(width); i++ {
		gc.DrawImage(source)
		gc.Translate(1, 0)
	}
	gc.Close()

	ximg := xgraphics.NewConvert(driver.util, dest)

	ximg.CreatePixmap()
	ximg.XDraw()
	ximg.XExpPaint(win, 0, 0)

}