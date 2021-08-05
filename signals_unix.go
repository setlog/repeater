// +build linux darwin

package repeater

import (
	"os"
	"syscall"
)

func cancellationSignals() []os.Signal {
	return []os.Signal{os.Interrupt, syscall.SIGTERM}
}
