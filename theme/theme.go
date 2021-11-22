package theme

import (
	"fmt"
	"path/filepath"
	"runtime"
)

type Theme struct {
	name string
}

func NewTheme(name string) *Theme {
	return &Theme{
		name: name,
	}
}

func (t *Theme) GetTitleBarImagePath() string {
	_, b, _, _ := runtime.Caller(0)
	basepath   := filepath.Dir(b)

	return fmt.Sprintf("%s/../resources/theme/%s/titlebar.png", basepath, t.name)
}