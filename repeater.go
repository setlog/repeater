package repeater

import (
	"context"
	"os/signal"
	"time"
)

// Repeater implements a periodic timer which supports context-cancellation.
type Repeater struct {
	processor Processor
}

// New constructs and returns a new *Repeater to periodically invoke given Processor.
func New(processor Processor) *Repeater {
	if processor == nil {
		panic("processor must be non-nil")
	}
	return &Repeater{
		processor: processor,
	}
}

// Run calls processor.Process() every time the duration specified by interval has passed and blocks
// until the given context is cancelled, the application receives an interrupt signal or an invocation
// of processor.Process() panics.
//
// If makeFirstCallImmediately is true, the first invocation of processor.Process() will happen immediately
// instead of after the first interval has passed.
func (r *Repeater) Run(parentContext context.Context, interval time.Duration, makeFirstCallImmediately bool) {
	if interval <= 0 {
		panic("interval must be > 0")
	}

	ctx, cancelFunc := signal.NotifyContext(parentContext, cancellationSignals()...)
	defer cancelFunc()

	defer r.processor.CleanUp()

	r.repeat(ctx, interval, makeFirstCallImmediately)
}

func (r *Repeater) repeat(ctx context.Context, interval time.Duration, makeFirstCallImmediately bool) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	if makeFirstCallImmediately {
		r.processor.Process(ctx)
	}
	for {
		select {
		case <-ticker.C:
			r.processor.Process(ctx)
		case <-ctx.Done():
			return
		}
	}
}
