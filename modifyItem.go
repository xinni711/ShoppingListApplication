package main

import (
	"fmt"
	"strings"
)

func (sl slc) modifyItem() {

	var item string
	var newItem string = ""
	var cat1 string
	var qty int
	var unit int
	var index int
	i := 0

	fmt.Println("Modify Items.\nWhich item would you wish to modify?")
	fmt.Scanln(&item)
	for k := range sl {
		if k == item {
			fmt.Printf("Current item name is %v - Category is %v - Quantity is %v - Unit Cost is %v\n", k, cat[sl[item].category], sl[item].quantity, sl[item].unitCost)
		} else {
			i++
		}
	}

	if i == len(sl) {
		fmt.Println("Item Not Found!")
	} else {
		fmt.Println("Enter new name. Enter for no change")
		fmt.Scanln(&newItem)
		fmt.Println("Enter new category. Enter for no change")
		fmt.Scanln(&cat1)

		if cat1 != "" {
			index = onlyExistingCategory(cat1) //check if existing category was selected, refer the func in addItems.go
		}

		fmt.Println("Enter new quantity. Enter for no change")
		fmt.Scanln(&qty)
		fmt.Println("Enter new unit cost. Enter for no change")
		fmt.Scanln(&unit)

		temp := sl[item]

		if cat1 == "" || strings.EqualFold(cat1, cat[index]) {
			fmt.Println("No changes to category made.")
		} else {
			temp.category = index
		}

		if qty == 0 {
			fmt.Println("No changes to quantity made.")
		} else {
			temp.quantity = qty
		}

		if unit == 0 {
			fmt.Println("No changes to unit cost made.")

		} else {
			temp.unitCost = unit
		}

		sl[item] = temp

		if newItem == "" || newItem == item {
			fmt.Println("No changes to item name made.")
		} else {
			sl[newItem] = sl[item]
			delete(sl, item)
		}

	}

}
