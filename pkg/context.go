package pkg

import (
	"io"

	"github.com/spf13/afero"
)

type AppContext struct {
	Name    string
	Version string

	Fs  afero.Fs
	Out io.Writer

	required struct{}
}
