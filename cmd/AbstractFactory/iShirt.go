package main

type IShirt interface {
	SetLogo(logo string)
	GetLogo() string
	SetSize(size string)
	GetSize() string
}
