package main

type IGun interface {
	GetName() string
	SetName(string)
	GetPower() int
	SetPower(int)
}
