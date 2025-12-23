package projectroot

import (
	"path/filepath"
	"runtime"
)

var (
	Path string
)

func init() {
	_, f, _, ok := runtime.Caller(0)
	if !ok {
		panic("could not determine project root path")
	}

	Path = filepath.Join(filepath.Dir(f), "../..")
}
