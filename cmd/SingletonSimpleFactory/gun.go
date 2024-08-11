package main

type Gun struct {
	name  string
	power int
}

func (gun *Gun) GetName() string {
	return gun.name
}

func (gun *Gun) SetName(name string) {
	gun.name = name
}

func (gun *Gun) GetPower() int {
	return gun.power
}

func (gun *Gun) SetPower(power int) {
	gun.power = power
}
