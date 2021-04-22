package main

import "fmt"

func (sl slc) printShoppingList() {

	fmt.Println("Shopping List Contents:")

	for k := range sl {
		fmt.Println("Category: ", cat[sl[k].category], " - Item: ", k, " Quantity: ", sl[k].quantity, " Unit Cost: ", sl[k].unitCost)
	}

}
