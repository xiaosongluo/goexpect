package demo

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ExampleTestSuite struct {
	suite.Suite
}

// @Setup( func = "",paras="[]" )
// @Name( name = "加法与减法测试")
func (suite *ExampleTestSuite) TestAddAndMinus() {
	assert.Equal(suite.T(), 5, 1+4)
	assert.Equal(suite.T(), 3, 4-1)
}

func (suite *ExampleTestSuite) TestMultiply() {
	suite.Assert()
}

func (suite *ExampleTestSuite) TestDivide() {
}

// In order for 'go test' to run this suite, we need to create a normal test function and pass our suite to suite.Run
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}
