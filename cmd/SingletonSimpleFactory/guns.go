package main

type Ak47 struct {
	Gun
}

type Musket struct {
	Gun
}

func newAk47(name string, power int) IGun {
	return &Ak47{
		Gun: Gun{
			name:  name,
			power: power,
		},
	}
}

func newMusket(name string, power int) IGun {
	return &Musket{
		Gun: Gun{
			name:  name,
			power: power,
		},
	}
}
