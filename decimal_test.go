package cnos

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type decimalSuite struct {
	suite.Suite
}

func TestDecimalSuite(t *testing.T) {
	suite.Run(t, new(decimalSuite))
}

func (d *decimalSuite) TestDecimal() {
	t := Decimal{Base: 94188}
	r1, err := t.GetDecimal()
	d.NoError(err)
	d.Equal("94188", r1)
	r2, err := t.GetBinary()
	d.NoError(err)
	d.Equal("00010110111111101100", r2)
	r3, err := t.GetHex()
	d.NoError(err)
	d.Equal("16fec", r3)
}
