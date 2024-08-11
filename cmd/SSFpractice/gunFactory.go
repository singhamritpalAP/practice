package main

import (
	"fmt"
	"sync"
)

var ak47, musket IGun
var lock sync.Mutex

func GetAk47() IGun {
	if ak47 == nil {
		lock.Lock()
		defer lock.Unlock()
		if ak47 == nil {
			fmt.Println("Creating new ak47 instance")
			ak47 = NewAk47()
		}
	}
	return ak47

}

func GetMusket() IGun {
	if musket == nil {
		lock.Lock()
		defer lock.Unlock()
		if musket == nil {
			fmt.Println("Creating new musket instance")
			musket = NewAk47()
		}
	}
	return musket
}
