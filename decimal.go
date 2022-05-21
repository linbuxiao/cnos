package cnos

import (
	"strconv"
)

type Decimal struct {
	Base int
}

func (d Decimal) GetBinary() (string, error) {
	t := d.Base
	var arr []int
	for t > 0 {
		arr = append(arr, t%2)
		t = t / 2
	}
	result := ""
	for i := len(arr) - 1; i >= 0; i-- {
		result += strconv.Itoa(arr[i])
	}
	return fixZero(result), nil
}

func (d Decimal) GetDecimal() (string, error) {
	return strconv.Itoa(d.Base), nil
}

func (d Decimal) GetHex() (string, error) {
	t := d.Base
	m := GetDecimal2HexMap()
	var arr []string
	for t > 0 {
		x := t % 16
		arr = append(arr, m[x])
		t = t / 16
	}
	result := ""
	for i := len(arr) - 1; i >= 0; i-- {
		result += arr[i]
	}
	return result, nil
}
