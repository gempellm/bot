package product

import "fmt"

var allProducts = []Product{
	{Title: "one"},
	{Title: "two"},
	{Title: "three"},
	{Title: "four"},
	{Title: "five"},
}

type Product struct {
	Title string
}

func (p *Product) String() string {
	return fmt.Sprintf("Product{Title: %v}", p.Title)
}
