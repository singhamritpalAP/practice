package main

type ISportsFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}
