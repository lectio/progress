package progress

import (
	"context"
	"io"
)

// ReaderProgressReporter is sent to this package's methods if activity progress reporting is expected for an io.Reader
type ReaderProgressReporter interface {
	StartReportableReaderActivityInBytes(ctx context.Context, summary string, exepectedBytes int64, inputReader io.Reader) io.Reader
	CompleteReportableActivityProgress(ctx context.Context, summary string)
}

// BoundedProgressReporter is one observation method for live reporting of long-running processes where the upper bound is known
type BoundedProgressReporter interface {
	StartReportableActivity(ctx context.Context, summary string, expectedItems int)
	IncrementReportableActivityProgress(ctx context.Context, incrementBy int)
	CompleteReportableActivityProgress(ctx context.Context, summary string)
}

// ExceptionCollector collect errors and warnings for later printing by a reporter
type ExceptionCollector interface {
	CollectError(context.Context, error) bool
	MaxErrorsCollected(context.Context) bool
	CollectWarning(ctx context.Context, code, message string) bool
}
