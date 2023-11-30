package main

import "fmt"

type ISportsFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

type IShoe interface {
}

type IShirt interface {
}

func GetSportsFactory(brand string) (ISportsFactory, error) {
	if brand == "adidas" {
		return &Adidas{}, nil
	}

	if brand == "nike" {
		return &Nike{}, nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}

type Adidas struct {
}

func (a *Adidas) makeShoe() IShoe {
	//TODO implement me
	return &AdidasShoe{
		Shoe{
			Size: 32,
			Logo: "chort",
		},
	}
}

func (a *Adidas) makeShirt() IShirt {
	return &AdidasShort{
		Shirt{
			Logo: "hello adidas",
			Size: 44,
		},
	}
}

type Shoe struct {
	Logo string
	Size int
}

type Shirt struct {
	Logo string
	Size int
}

type AdidasShoe struct {
	Shoe
}

type AdidasShort struct {
	Shirt
}

type NikeShoe struct {
	Shoe
}

type NikeShort struct {
	Shirt
}

type Nike struct {
}

func (n *Nike) makeShoe() IShoe {
	return &NikeShoe{
		Shoe{
			Size: 43,
			Logo: "re",
		},
	}
}

func (n *Nike) makeShirt() IShirt {
	return &NikeShort{
		Shirt{
			Size: 43,
			Logo: "Do",
		},
	}
}

func main() {
	adidasFactory, _ := GetSportsFactory("adidas")
	nikeFactory, _ := GetSportsFactory("nike")

	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()

	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()

}
