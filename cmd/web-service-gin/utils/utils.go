package utils

import "strconv"

type Utils struct {
}

func (utils *Utils) GetIntegerValueFromString(str string) (int, error) {
	intValue, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return intValue, nil
}
