package cnos

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type hexSuite struct {
	suite.Suite
}

func TestHexSuite(t *testing.T) {
	suite.Run(t, new(hexSuite))
}

func (h *hexSuite) TestHex() {
	t := Hex{Base: "16fec"}
	r1, err := t.GetHex()
	h.NoError(err)
	h.Equal("16fec", r1)
	r2, err := t.GetDecimal()
	h.NoError(err)
	h.Equal("94188", r2)
	r3, err := t.GetBinary()
	h.NoError(err)
	h.Equal("00010110111111101100", r3)
}
