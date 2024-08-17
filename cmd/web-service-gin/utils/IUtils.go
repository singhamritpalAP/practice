package utils

type IUtils interface {
	GetIntegerValueFromString(str string) (int, error)
}
