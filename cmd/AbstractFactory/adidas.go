package main

type Adidas struct {
}

func GetAdidas() ISportsFactory {
	return &Adidas{}
}

func (adidas *Adidas) makeShoe() IShoe {
	return &AdidasShoes{
		Shoes{
			logo: "",
			size: "",
		},
	}
}

func (adidas *Adidas) makeShirt() IShirt {
	return &AdidasShirts{
		Shirt{
			logo: "",
			size: "",
		},
	}
}
