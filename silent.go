package progress

import (
	"context"
	"io"
)

type SilentProgressReporter struct{}

func (pr SilentProgressReporter) StartReportableActivity(ctx context.Context, summary string, expectedItems int) {
}

func (pr SilentProgressReporter) StartReportableReaderActivityInBytes(ctx context.Context, summary string, exepectedBytes int64, inputReader io.Reader) io.Reader {
	return inputReader
}

func (pr SilentProgressReporter) IncrementReportableActivityProgress(ctx context.Context, incrementBy int) {
}

func (pr SilentProgressReporter) CompleteReportableActivityProgress(ctx context.Context, summary string) {
}

func (pr SilentProgressReporter) CollectError(context.Context, error) bool {
	return true
}

func (pr SilentProgressReporter) MaxErrorsCollected(context.Context) bool {
	return false
}

func (pr SilentProgressReporter) CollectWarning(ctx context.Context, code, message string) bool {
	return true
}
