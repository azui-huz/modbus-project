package cycreader

import (
	"context"
	"time"
)

type CyclicReader struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func StartPeriodic(ctx context.Context, interval time.Duration, readFn func()) *CyclicReader {
	cctx, cancel := context.WithCancel(ctx)
	cr := &CyclicReader{ctx: cctx, cancel: cancel}
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			select {
			case <-cctx.Done():
				return
			case <-ticker.C:
				readFn()
			}
		}
	}()
	return cr
}

func (c *CyclicReader) Stop() { c.cancel() }
