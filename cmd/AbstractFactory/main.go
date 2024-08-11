package main

import "fmt"

func main() {
	adidasFactory := GetAdidas()
	nikeFactory := GetNike()

	adidasShoe := adidasFactory.makeShoe()
	adidasShoe.SetLogo("Adidas")
	adidasShoe.SetSize("7")

	adidasShirt := adidasFactory.makeShirt()
	adidasShirt.SetLogo("Adidas")
	adidasShirt.SetSize("M")

	nikeShoe := nikeFactory.makeShoe()
	nikeShoe.SetLogo("Nike")
	nikeShoe.SetSize("8")

	nikeShirt := nikeFactory.makeShirt()
	nikeShirt.SetLogo("Nike")
	nikeShirt.SetSize("S")

	printShirtDetails(nikeShirt)
	printShoeDetails(nikeShoe)
	printShirtDetails(adidasShoe)
	printShoeDetails(adidasShirt)

}

func printShirtDetails(shoe IShoe) {
	fmt.Println("Shirt brand: ", shoe.GetLogo())
	fmt.Println("Shirt size: ", shoe.GetSize())
}

func printShoeDetails(shirt IShirt) {
	fmt.Println("Shoe brand: ", shirt.GetLogo())
	fmt.Println("Shoe size: ", shirt.GetSize())
}
