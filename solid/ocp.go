package main

import "fmt"

type Color int

const (
	red Color = iota + 1
	green
	blue
)

type Size int

const (
	large = iota + 1
	medium
	small
)

type Product struct {
	name  string
	color Color
	size  Size
}

type BadFilter struct{}

func (f *BadFilter) filterByColor(prods []Product, color Color) []Product {
	result := make([]Product, 0)
	for _, prod := range prods {
		if prod.color == color {
			result = append(result, prod)
		}
	}
	return result
}

func (f *BadFilter) filterBySize(prods []Product, size Size) []Product {
	result := make([]Product, 0)
	for _, prod := range prods {
		if prod.size == size {
			result = append(result, prod)
		}
	}
	return result
}

// ---
// define the interface for various spec

type Specification interface {
	IsSatisfied(prod *Product) bool
}

type ColorSpec struct {
	color Color
}

func (f ColorSpec) IsSatisfied(prod *Product) bool {
	return prod.color == f.color
}

type SizeSpec struct {
	size Size
}

func (f SizeSpec) IsSatisfied(prod *Product) bool {
	return prod.size == f.size
}

type MultiSpec struct {
	specs []Specification
}

func (f MultiSpec) IsSatisfied(prod *Product) bool {
	for _, spec := range f.specs {
		if !spec.IsSatisfied(prod) {
			return false
		}
	}
	return true
}

type BetterFilter struct{}

func (f *BetterFilter) Filter(prods []Product, spec Specification) []Product {
	result := make([]Product, 0)

	for _, prod := range prods {
		if spec.IsSatisfied(&prod) {
			result = append(result, prod)
		}
	}

	return result
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}
	products := []Product{apple, tree, house}

	fmt.Println("Bad:")
	f := BadFilter{}
	for _, v := range f.filterByColor(products, green) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	fmt.Println("Better:")
	greenSpec := ColorSpec{green}
	bf := BetterFilter{}
	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	largeSpec := SizeSpec{large}
	largeGreenSpec := MultiSpec{[]Specification{largeSpec, greenSpec}}
	fmt.Print("Use MultiSpec to find large blue items:\n")
	for _, v := range bf.Filter(products, largeGreenSpec) {
		fmt.Printf(" - %s is large and green\n", v.name)
	}
}
