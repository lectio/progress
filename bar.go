package progress

import (
	"context"
	"fmt"
	"gopkg.in/cheggaaa/pb.v1"
	"io"
	"sync"
)

// BarReporter shows a horizontal progress bar on StdOut which is incremented as activities are completed
type BarReporter struct {
	prefix    string
	mu        sync.RWMutex
	bar       *pb.ProgressBar
	errList   []error
	warnList  []struct{ code, message string }
	maxErrors int
}

func (pr *BarReporter) StartReportableActivity(ctx context.Context, summary string, expectedItems int) {
	fmt.Printf("%s%s\n", pr.prefix, summary)
	pr.bar = pb.StartNew(expectedItems)
	pr.bar.ShowCounters = true
}

func (pr *BarReporter) StartReportableReaderActivityInBytes(ctx context.Context, summary string, exepectedBytes int64, inputReader io.Reader) io.Reader {
	pr.bar = pb.New(int(exepectedBytes)).SetUnits(pb.U_BYTES)
	pr.bar.Start()
	return pr.bar.NewProxyReader(inputReader)
}

func (pr *BarReporter) IncrementReportableActivityProgress(ctx context.Context) {
	pr.bar.Increment()
}

func (pr *BarReporter) IncrementReportableActivityProgressBy(ctx context.Context, incrementBy int) {
	pr.bar.Add(incrementBy)
}

func (pr *BarReporter) CompleteReportableActivityProgress(ctx context.Context, summary string) {
	pr.bar.FinishPrint(fmt.Sprintf("%s%s\n", pr.prefix, summary))

	if len(pr.errList) > 0 {
		for _, err := range pr.errList {
			fmt.Printf("%s%v\n", pr.prefix, err.Error())
		}
	}

	if len(pr.warnList) > 0 {
		for _, warning := range pr.warnList {
			fmt.Printf("%s%s %s\n", pr.prefix, warning.code, warning.message)
		}
	}
}

func (pr *BarReporter) CollectError(ctx context.Context, err error) bool {
	pr.mu.Lock()
	pr.errList = append(pr.errList, err)
	pr.mu.Unlock()
	return !pr.MaxErrorsCollected(ctx)
}

func (pr *BarReporter) MaxErrorsCollected(context.Context) bool {
	pr.mu.RLock()
	defer pr.mu.RUnlock()
	return pr.maxErrors > 0 && len(pr.errList) > pr.maxErrors
}

func (pr *BarReporter) CollectWarning(ctx context.Context, code, message string) bool {
	pr.mu.Lock()
	pr.warnList = append(pr.warnList, struct{ code, message string }{code: code, message: message})
	pr.mu.Unlock()
	return true
}
