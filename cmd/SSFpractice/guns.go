package main

type Ak47 struct {
	Gun
}

type Musket struct {
	Gun
}

func NewAk47() IGun {
	return &Ak47{
		Gun{
			name:  "",
			power: 0,
		},
	}
}

func NewMusket() IGun {
	return &Musket{
		Gun{
			name:  "",
			power: 0,
		},
	}
}
