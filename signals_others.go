// +build !linux,!darwin

package repeater

import (
	"os"
)

func cancellationSignals() []os.Signal {
	return []os.Signal{os.Interrupt}
}
