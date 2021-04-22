package main

import (
	"fmt"
	"strings"
)

func addCategory() {

	var newCategory string
	var i int

	fmt.Println("Add New Category Name")
	fmt.Println("What is the New Category Name to add?")
	fmt.Scanln(&newCategory)

	for i = 0; i < (len(cat)); i++ {
		if newCategory == "" {
			fmt.Println("No Input Found!\n")
			break
		} else if strings.EqualFold(newCategory, cat[i]) {
			fmt.Printf("Category: %v already exist at index %v\n", strings.Title(newCategory), i)
			break
		}
	}

	if i == len(cat) {
		cat = append(cat, strings.Title(newCategory))
		fmt.Printf("New Category: %v added at index %v\n", cat[i], i)
	}
}
