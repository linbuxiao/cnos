package cnos

import "strconv"

type Hex struct {
	Base string
}

func (h Hex) GetBinary() (string, error) {
	m := GetHex2BinaryMap()
	result := ""
	for i := 0; i < len(h.Base); i++ {
		result += m[string(h.Base[i])]
	}
	return result, nil
}

func (h Hex) GetDecimal() (string, error) {
	m := GetHex2DecimalMap()
	result := 0
	for i := 0; i < len(h.Base); i++ {
		j := len(h.Base) - (i + 1)
		y, err := powerOf16(j)
		if err != nil {
			return "", err
		}
		x := m[string(h.Base[i])]
		result += x * y
	}
	return strconv.Itoa(result), nil
}

func (h Hex) GetHex() (string, error) {
	return h.Base, nil
}
