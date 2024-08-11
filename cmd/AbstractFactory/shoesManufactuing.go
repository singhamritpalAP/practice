package main

type Shoes struct {
	logo string
	size string
}

func (shoe *Shoes) SetLogo(logo string) {
	shoe.logo = logo
}

func (shoe *Shoes) SetSize(size string) {
	shoe.size = size
}

func (shoe *Shoes) GetLogo() string {
	return shoe.logo
}

func (shoe *Shoes) GetSize() string {
	return shoe.size
}
