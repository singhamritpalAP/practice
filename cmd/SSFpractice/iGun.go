package main

type IGun interface {
	SetNameP(name string)
	GetNameP() string
	SetPowerP(power int)
	GetPowerP() int
}
