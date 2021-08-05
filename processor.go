package repeater

import "context"

// Processor is the interface for all types which can be invoked by a *Repeater.
type Processor interface {
	// Process is called periodically by a *Repeater.
	Process(ctx context.Context)

	// CleanUp is called when the periodic call-cycle is stopped due to either context-cancellation or Process() having panicked.
	CleanUp()
}
