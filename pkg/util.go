package pkg

import (
	"fmt"
	"io"
)

func PrintVersion(out io.Writer, version string) {
	_, _ = fmt.Fprintf(out, "%s\n", version)
}
