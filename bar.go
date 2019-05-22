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
	Prefix    string
	ErrList   []error
	WarnList  []struct{ code, message string }
	MaxErrors int

	mu  sync.RWMutex
	bar *pb.ProgressBar
}

func NewBarReporter(prefix string) *BarReporter {
	result := new(BarReporter)
	result.Prefix = prefix
	return result
}

func (pr *BarReporter) StartReportableActivity(ctx context.Context, summary string, expectedItems int) {
	fmt.Printf("%s%s\n", pr.Prefix, summary)
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
	pr.bar.FinishPrint(fmt.Sprintf("%s%s\n", pr.Prefix, summary))

	if len(pr.ErrList) > 0 {
		for _, err := range pr.ErrList {
			fmt.Printf("%s%v\n", pr.Prefix, err.Error())
		}
	}

	if len(pr.WarnList) > 0 {
		for _, warning := range pr.WarnList {
			fmt.Printf("%s%s %s\n", pr.Prefix, warning.code, warning.message)
		}
	}
}

func (pr *BarReporter) CollectError(ctx context.Context, err error) bool {
	pr.mu.Lock()
	pr.ErrList = append(pr.ErrList, err)
	pr.mu.Unlock()
	return !pr.MaxErrorsCollected(ctx)
}

func (pr *BarReporter) MaxErrorsCollected(context.Context) bool {
	pr.mu.RLock()
	defer pr.mu.RUnlock()
	return pr.MaxErrors > 0 && len(pr.ErrList) > pr.MaxErrors
}

func (pr *BarReporter) CollectWarning(ctx context.Context, code, message string) bool {
	pr.mu.Lock()
	pr.WarnList = append(pr.WarnList, struct{ code, message string }{code: code, message: message})
	pr.mu.Unlock()
	return true
}
