package main

import (
	"fmt"
	"sync"
)

var singleTonAk47 IGun
var lock sync.Mutex

func GetAk47(name string, power int) IGun {
	if singleTonAk47 == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleTonAk47 == nil {
			fmt.Println("Creating new instance for: ", name)
			singleTonAk47 = newAk47(name, power)
		} else {
			fmt.Println("Instance already exists not creating for ", name)
		}
	} else {
		fmt.Println("Instance already exists not creating for ", name)
	}
	return singleTonAk47
}
func GetMusket(name string, power int) IGun {
	return newMusket(name, power)
}
