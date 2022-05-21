package cnos

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type binarySuite struct {
	suite.Suite
}

func TestBinarySuite(t *testing.T) {
	suite.Run(t, new(binarySuite))
}

func (b *binarySuite) TestBinary() {
	base := "1100"
	t := Binary{Base: base}
	r1, err := t.GetBinary()
	b.NoError(err)
	b.Equal(base, r1)
	r2, err := t.GetDecimal()
	b.NoError(err)
	b.Equal("12", r2)
	r3, err := t.GetHex()
	b.NoError(err)
	b.Equal("c", r3)
}
