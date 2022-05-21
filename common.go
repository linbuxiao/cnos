package cnos

import (
	"errors"
	"strconv"
	"strings"
)

var ErrUnsupportedBase = errors.New("unsupported base")
var ErrUnsupportedPower = errors.New("unsupported power")

var baseMap = map[int]SystemRepo{
	2:  Binary{},
	10: Decimal{},
	16: Hex{},
}

var binary2HexMap = map[string]string{
	"0000": "0",
	"0001": "1",
	"0010": "2",
	"0011": "3",
	"0100": "4",
	"0101": "5",
	"0110": "6",
	"0111": "7",
	"1000": "8",
	"1001": "9",
	"1010": "a",
	"1011": "b",
	"1100": "c",
	"1101": "d",
	"1110": "e",
	"1111": "f",
}

type SystemRepo interface {
	GetBinary() (string, error)
	GetDecimal() (string, error)
	GetHex() (string, error)
}

func NewCnos(n int, base string) (SystemRepo, error) {
	r := checkBase(n)
	if !r {
		return nil, ErrUnsupportedBase
	}
	var cnos SystemRepo
	switch n {
	case 2:
		cnos = Binary{Base: fixZero(base)}
	case 10:
		t, err := strconv.Atoi(base)
		if err != nil {
			return nil, err
		}
		cnos = Decimal{Base: t}
	case 16:
		cnos = Hex{Base: strings.ToLower(base)}
	default:
		return nil, nil
	}

	return cnos, nil
}

func checkBase(n int) bool {
	_, ok := baseMap[n]
	return ok
}

func GetValue(cn SystemRepo, to int) (string, error) {
	r := checkBase(to)
	if !r {
		return "", ErrUnsupportedBase
	}
	switch to {
	case 2:
		return cn.GetBinary()
	case 10:
		return cn.GetDecimal()
	case 16:
		return cn.GetHex()
	default:
		return "", nil
	}
}

func fixZero(s string) string {
	t := len(s) % 4
	if t == 0 {
		return s
	}
	t = 4 - t
	for t > 0 {
		s = "0" + s
		t--
	}

	return s
}

func powerOf2(index int) (int, error) {
	return powerOfN(2, index)
}

func powerOf16(index int) (int, error) {
	return powerOfN(16, index)
}

func powerOfN(base int, index int) (int, error) {
	if index < 0 {
		return 0, ErrUnsupportedPower
	}

	if index == 0 {
		return 1, nil
	}

	index--
	result := base

	for index > 0 {
		result *= base
		index--
	}

	return result, nil
}

func GetBinary2HexMap() map[string]string {
	return binary2HexMap
}

func GetHex2BinaryMap() map[string]string {
	m := make(map[string]string)
	for k, v := range binary2HexMap {
		m[v] = k
	}
	return m
}

func GetDecimal2HexMap() map[int]string {
	word := "abcdef"
	m := make(map[int]string)
	t := 15
	for t >= 0 {
		if t >= 10 {
			m[t] = string(word[t-10])
		} else {
			m[t] = strconv.Itoa(t)
		}
		t--
	}
	return m
}

func GetHex2DecimalMap() map[string]int {
	m := GetDecimal2HexMap()
	result := make(map[string]int)
	for k, v := range m {
		result[v] = k
	}
	return result
}

func cutBinary(s string) []string {
	var arr []string
	for i := 0; i < len(s); i += 4 {
		arr = append(arr, s[i:i+4])
	}
	return arr
}
