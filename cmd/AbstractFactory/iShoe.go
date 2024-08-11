package main

type IShoe interface {
	SetLogo(logo string)
	SetSize(size string)
	GetLogo() string
	GetSize() string
}
