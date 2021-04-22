package main

import (
	"fmt"
	"strings"
)

func modifyDeleteCategory() {

	var choice int

	fmt.Println("Modify or Delete Category.")
	fmt.Println("1. Modify Category")
	fmt.Println("2. Delete Category")
	fmt.Println("Select your choice ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		modifyCategory()
	case 2:
		slc(shoppingListContent).deleteCategory()
	default:
		fmt.Println("Please select 1 or 2")
	}
}

func modifyCategory() {
	var input string
	var newInput string

	fmt.Println("Enter category to modify: ")
	fmt.Scanln(&input)

	for i, v := range cat {
		if strings.EqualFold(v, input) {
			fmt.Println("Enter new category name ")
			fmt.Scanln(&newInput)
			_, result := ifExist(newInput)
			if result == true || newInput == "" {
				fmt.Println("Entry has the same name as existing category or entry is empty, not accepted")
			} else {
				cat[i] = strings.Title(newInput)
			}
			break
		} else {
			if i == len(cat)-1 {
				fmt.Println("Nothing to modify")
			}
		}
	}

}

func (sl slc) deleteCategory() {

	var input string

	fmt.Println("Enter category to delete: ")
	fmt.Scanln(&input)

	for i, v := range cat {
		if strings.EqualFold(v, input) {
			for k := range sl {
				if sl[k].category == i {
					delete(sl, k)
				} else if sl[k].category > i {
					temp := sl[k]
					temp.category = sl[k].category - 1
					sl[k] = temp
				}
			}
			cat = append(cat[:i], cat[i+1:]...)
			break

		} else {
			if i == len(cat)-1 {
				fmt.Println("Nothing to delete")
			}
		}
	}

}
