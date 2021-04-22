package main

import "fmt"

func (sl slc) deleteItem() {

	var itemToDelete string

	fmt.Println("Delete Item.")
	fmt.Println("Enter item name to delete: ")
	fmt.Scanln(&itemToDelete)

	_, ok := sl[itemToDelete]

	if ok == false {
		fmt.Println("Item not found.Nothing to delete!")
	} else {
		delete(sl, itemToDelete)
		fmt.Println("Deleted ", itemToDelete)
	}

}
