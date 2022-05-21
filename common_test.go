package cnos

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type commonSuite struct {
	suite.Suite
}

func TestCommonSuite(t *testing.T) {
	suite.Run(t, new(commonSuite))
}

func (c *commonSuite) Test_fixZero() {
	m := map[string]string{
		"00":                "0000",
		"10":                "0010",
		"1000":              "1000",
		"10110111111101100": "00010110111111101100",
	}
	for k, v := range m {
		c.Equal(v, fixZero(k))
	}
}

func (c *commonSuite) Test_powerOf2() {
	m := map[int]int{
		0: 1,
		2: 4,
		4: 16,
		8: 256,
	}

	for k, v := range m {
		x, err := powerOf2(k)
		c.NoError(err)
		c.Equal(v, x)
	}
}

func (c *commonSuite) Test_GetDecimal2HexMap() {
	m := map[int]string{
		10: "a",
		11: "b",
		12: "c",
		13: "d",
		14: "e",
		15: "f",
	}
	t := GetDecimal2HexMap()
	for k, v := range m {
		c.Equal(v, t[k])
	}
}
