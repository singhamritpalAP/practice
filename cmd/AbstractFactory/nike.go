package main

type Nike struct {
}

func GetNike() ISportsFactory {
	return &Nike{}
}

func (nike *Nike) makeShoe() IShoe {
	return &NikeShoes{
		Shoes{
			logo: "",
			size: "",
		},
	}
}

func (nike *Nike) makeShirt() IShirt {
	return &NikeShirt{
		Shirt{
			logo: "",
			size: "",
		},
	}
}
