package main

import "fmt"

func main() {
	ak47Model1 := GetAk47()
	ak47Model2 := GetAk47()
	musketModel1 := GetMusket()

	ak47Model1.SetNameP("AK47 model - 1")
	ak47Model1.SetPowerP(9)
	printDetails(ak47Model1)

	ak47Model2.SetNameP("AK47 model - 2")
	ak47Model2.SetPowerP(10)

	musketModel1.SetNameP("Musket model - 1")
	musketModel1.SetPowerP(7)

	printDetails(ak47Model1)
	printDetails(musketModel1)
}

func printDetails(gun IGun) {
	fmt.Println("Gun name: ", gun.GetNameP())
	fmt.Println("Gun power: ", gun.GetPowerP())
}
