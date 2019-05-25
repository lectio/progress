package progress

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ProgressSuite struct {
	suite.Suite
}

func (suite *ProgressSuite) SetupSuite() {
}

func (suite *ProgressSuite) TearDownSuite() {
}

func (suite *ProgressSuite) TestBarTypes() {
	br := NewBarReporter("")
	suite.Implements((*ReaderProgressReporter)(nil), br)
	suite.Implements((*BoundedProgressReporter)(nil), br)
	suite.Implements((*ExceptionCollector)(nil), br)
}

func (suite *ProgressSuite) TestSummaryTypes() {
	br := NewSummaryReporter("")
	suite.Implements((*ReaderProgressReporter)(nil), br)
	suite.Implements((*BoundedProgressReporter)(nil), br)
	suite.Implements((*ExceptionCollector)(nil), br)
}

func (suite *ProgressSuite) TestSilentTypes() {
	br := NewSilentReporter("")
	suite.Implements((*ReaderProgressReporter)(nil), br)
	suite.Implements((*BoundedProgressReporter)(nil), br)
	suite.Implements((*ExceptionCollector)(nil), br)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(ProgressSuite))
}
