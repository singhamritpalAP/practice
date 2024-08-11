package main

type Gun struct {
	name  string
	power int
}

func (gun *Gun) SetNameP(name string) {
	gun.name = name
}

func (gun *Gun) SetPowerP(power int) {
	gun.power = power
}

func (gun *Gun) GetNameP() string {
	return gun.name
}

func (gun *Gun) GetPowerP() int {
	return gun.power
}
