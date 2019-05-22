package progress

import (
	"context"
	"fmt"
	"io"
	"sync"
)

type SummaryReporter struct {
	writer         io.Writer
	prefix         string
	mu             sync.RWMutex
	errorsReported uint
	maxErrors      uint
}

func (pr *SummaryReporter) StartReportableActivity(ctx context.Context, summary string, expectedItems int) {
	fmt.Fprintf(pr.writer, "%s%s\n", pr.prefix, summary)
}

func (pr *SummaryReporter) StartReportableReaderActivityInBytes(ctx context.Context, summary string, exepectedBytes int64, inputReader io.Reader) io.Reader {
	fmt.Fprintf(pr.writer, "%s%s\n", pr.prefix, summary)
	return inputReader
}

func (pr *SummaryReporter) IncrementReportableActivityProgress(ctx context.Context, incrementBy int) {
}

func (pr *SummaryReporter) CompleteReportableActivityProgress(ctx context.Context, summary string) {
	fmt.Fprintf(pr.writer, "%s%s\n", pr.prefix, summary)
}

func (pr *SummaryReporter) CollectError(ctx context.Context, err error) bool {
	pr.mu.Lock()
	pr.errorsReported++
	pr.mu.Unlock()
	fmt.Fprintf(pr.writer, "%s%v\n", pr.prefix, err.Error())
	return !pr.MaxErrorsCollected(ctx)
}

func (pr *SummaryReporter) MaxErrorsCollected(context.Context) bool {
	pr.mu.RLock()
	defer pr.mu.RUnlock()
	return pr.maxErrors > 0 && pr.errorsReported > pr.maxErrors
}

func (pr *SummaryReporter) CollectWarning(ctx context.Context, code, message string) bool {
	fmt.Fprintf(pr.writer, "%s%s %s\n", pr.prefix, code, message)
	return true
}
