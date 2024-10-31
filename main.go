package main

import "fmt"

type Product struct {
	Name  string
	Price int
	Tax   int
}

func MoneyChanges(amount int, products []Product) []int {
	total := 0
	for _, product := range products {
		total += product.Price + product.Tax
	}

	change := amount - total
	if change <= 0 {
		return []int{}
	}

	denominations := []int{1000, 500, 200, 100, 50, 20, 10, 5, 1}
	var result []int

	for _, denomination := range denominations {
		for change >= denomination {
			result = append(result, denomination)
			change -= denomination
		}
	}

	return result
}

func main() {
	products := []Product{
		{Name: "Baju", Price: 5000, Tax: 500},
		{Name: "Celana", Price: 3000, Tax: 300},
	}

	change := MoneyChanges(10000, products)
	fmt.Println(change) // Output: [1000 200]

	changeEmpty := MoneyChanges(0, []Product{})
	fmt.Println(changeEmpty) // Output: []

	changeSame := MoneyChanges(10000, products) // Total = 5800
	fmt.Println(changeSame)                     // Output: []
}
