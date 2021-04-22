package main

import (
	"fmt"
)

func generateReport() {

	var choice int

	fmt.Println("Generate Report")
	fmt.Println("1. Total Cost of each category")
	fmt.Println("2. List of item by category")
	fmt.Println("3. Main Menu\n")
	fmt.Println("Choose your report: ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		slc(shoppingListContent).totalCost()
	case 2:
		fmt.Println("List by Category:")
		slc(shoppingListContent).listOfItem()
	case 3:
		printMenu() //print the main shopping list menu
	default:
		fmt.Println("Please select 1 to 3.")
	}
}

func (sl slc) totalCost() {

	sum := make([]int, len(cat))

	for k := range sl {
		for i := 0; i < len(cat); i++ {
			if sl[k].category == i {
				sum[i] += sl[k].quantity * sl[k].unitCost
			}
		}
	}
	fmt.Println("Total cost by Category.")
	for i := 0; i < len(cat); i++ {
		fmt.Println(cat[i], "cost :", sum[i])
	}

}

func (sl slc) listOfItem() {

	for i := 0; i < len(cat); i++ {
		for k := range sl {
			if cat[i] == cat[sl[k].category] {
				fmt.Println("Category: ", cat[sl[k].category], " - Item: ", k, " Quantity: ", sl[k].quantity, " Unit Cost: ", sl[k].unitCost)
			}
		}

	}
}
