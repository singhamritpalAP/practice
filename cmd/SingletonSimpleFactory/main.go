package main

import "fmt"

func main() {
	ak47 := GetAk47("AK47 Model 1", 10)
	ak472 := GetAk47("AK47 Model 2", 10)
	musket := GetMusket("Musket Model 1", 10)

	printGunDetails(ak47)
	printGunDetails(ak472)
	printGunDetails(musket)
	musket.SetPower(8)
	fmt.Println("Power updated")
	printGunDetails(musket)

}

func printGunDetails(gun IGun) {
	fmt.Println("Gun Name: ", gun.GetName())
	fmt.Println("Gun Power: ", gun.GetPower())
}
