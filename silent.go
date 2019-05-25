package progress

import (
	"context"
	"io"
)

type SilentReporter struct{}

func NewSilentReporter(prefix string) *SilentReporter {
	return &SilentReporter{}
}

func (pr SilentReporter) StartReportableActivity(ctx context.Context, summary string, expectedItems int) {
}

func (pr SilentReporter) StartReportableReaderActivityInBytes(ctx context.Context, summary string, exepectedBytes int64, inputReader io.Reader) io.Reader {
	return inputReader
}

func (pr SilentReporter) IncrementReportableActivityProgress(ctx context.Context, incrementBy int) {
}

func (pr SilentReporter) CompleteReportableActivityProgress(ctx context.Context, summary string) {
}

func (pr SilentReporter) CollectError(context.Context, error) bool {
	return true
}

func (pr SilentReporter) MaxErrorsCollected(context.Context) bool {
	return false
}

func (pr SilentReporter) CollectWarning(ctx context.Context, code, message string) bool {
	return true
}
