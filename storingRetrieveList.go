package main

import (
	"fmt"
)

var shopList [][]string

func storingRetrieveList() {

	var choice int

	fmt.Println("Store or Retrieve Shopping List")
	fmt.Println("Current stored shopping list is ", len(shopList))
	fmt.Println("1. Save Shopping List")
	fmt.Println("2. Retrieve Shopping List")
	fmt.Println("Select your choice ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		slc(shoppingListContent).storeList()
	case 2:
		retrieveList()
	default:
		fmt.Println("Please select 1 or 2")
	}

}

func (sl slc) storeList() {

	var list []string
	var item string

	choice := 0

	for choice != 2 {

		fmt.Println("Press 1 to Add Item.\nPress 2 to Complete and Save")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("Enter item: ")
			fmt.Scanln(&item)

			list = append(list, item)

		case 2:
			shopList = append(shopList, list)
			item = ""
			fmt.Println("Shopping list store in index ", len(shopList)-1)

			slc(shoppingListContent).ifExistingItem()

		default:
			fmt.Println("Please select 1 or 2")
		}
	}

}

func retrieveList() {

	var index int

	fmt.Println("Enter shopping list index to retrieve: ")
	fmt.Scanln(&index)

	if len(shopList) != 0 {
		if index >= len(shopList) {
			fmt.Println("No shopping list with index ", index)
		} else {
			fmt.Println("Shopping list index", index)
			slc(shoppingListContent).ifExistingItem()

		}
	} else {
		fmt.Println("Nothing is stored!")
	}

}

func (sl slc) ifExistingItem() {

	for i := range shopList[len(shopList)-1] {
		j := 0
		for k := range sl {
			if shopList[len(shopList)-1][i] == k {
				//fmt.Print(i+1, ".", k, "\n")
				fmt.Println(i+1, ".", k, "- Category: ", cat[sl[k].category], " Quantity: ", sl[k].quantity, " Unit Cost: ", sl[k].unitCost)
			} else {
				j++
			}
		}
		if j == len(sl) {
			fmt.Print(i+1, " . ", shopList[len(shopList)-1][i], "(Item not available!)\n")
		}
	}
}
