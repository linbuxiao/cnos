package cnos

import "strconv"

type Binary struct {
	Base string
}

func (b Binary) GetBinary() (string, error) {
	return b.Base, nil
}

func (b Binary) GetDecimal() (string, error) {
	p := 0
	for i, j := 0, len(b.Base)-1; i < len(b.Base) && j >= 0; i, j = i+1, j-1 {
		x, err := powerOf2(j)
		if err != nil {
			return "", err
		}

		y0 := string(b.Base[i])
		y1, err := strconv.Atoi(y0)
		if err != nil {
			return "", err
		}
		p += x * y1
	}

	return strconv.Itoa(p), nil
}

func (b Binary) GetHex() (string, error) {
	m := GetBinary2HexMap()
	arr := cutBinary(b.Base)
	result := ""
	for _, v := range arr {
		result += m[v]
	}
	return result, nil
}
