package main

type Shirt struct {
	logo string
	size string
}

func (shirt *Shirt) SetLogo(logo string) {
	shirt.logo = logo
}

func (shirt *Shirt) SetSize(size string) {
	shirt.size = size
}

func (shirt *Shirt) GetLogo() string {
	return shirt.logo
}

func (shirt *Shirt) GetSize() string {
	return shirt.size
}
