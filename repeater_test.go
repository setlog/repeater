package repeater_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/setlog/repeater"
)

var errBail error = fmt.Errorf("bail")

type testProcessor struct {
	callCount    int
	panicCount   int
	cleanupCount int
}

func (tp *testProcessor) Process(ctx context.Context) {
	tp.callCount++
	if tp.callCount == tp.panicCount {
		panic(errBail)
	}
}

func (tp *testProcessor) CleanUp() {
	if tp.callCount != tp.panicCount {
		tp.cleanupCount = -1
	} else {
		tp.cleanupCount++
	}
}

func TestCallNTimes(t *testing.T) {
	for targetCount := 1; targetCount < 3; targetCount++ {
		callNTimes(t, targetCount, false, false)
		callNTimes(t, targetCount, true, false)
		callNTimes(t, targetCount, false, true)
		callNTimes(t, targetCount, true, true)
	}
}

func callNTimes(t *testing.T, targetCount int, makeFirstCallImmediately bool, waitFull bool) {
	tp := &testProcessor{panicCount: targetCount}
	rep := repeater.New(tp)
	rep.WaitFull = waitFull
	defer func() {
		r := recover()
		if r == nil {
			t.Fatalf("targetCount = %d: tp did not panic", targetCount)
		}
		if r != errBail {
			t.Fatalf("targetCount = %d: unexpected panic value %v", targetCount, r)
		}
		if tp.callCount != targetCount {
			t.Fatalf("targetCount = %d: unexpected call count %d", targetCount, tp.callCount)
		}
		if tp.cleanupCount != 1 {
			t.Fatalf("targetCount = %d: unexpected cleanup count %d", targetCount, tp.cleanupCount)
		}
	}()
	rep.Run(context.TODO(), time.Nanosecond, makeFirstCallImmediately)
}
