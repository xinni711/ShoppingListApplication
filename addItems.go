package main

import (
	"fmt"
	"strings"
)

func (sl slc) addItem() {

	var item string
	var cat1 string
	var qty int
	var unit int
	var notEqual bool

	fmt.Println("What is the name of your item?")
	fmt.Scanln(&item)

	for notEqual != true {
		i := 0
		for k := range sl {
			if strings.EqualFold(k, item) {
				fmt.Println("Item exist! Try again!")
				fmt.Scanln(&item)
			} else {
				i++
			}
		}
		for i == len(sl) {
			notEqual = true
			break
		}
	}

	if item != "" { //Only accept non empty entry!

		fmt.Println("What category does it belong to?")
		fmt.Scanln(&cat1)

		index := onlyExistingCategory(cat1)

		fmt.Println("How many units are there?")
		fmt.Scanln(&qty)
		fmt.Println("How much does it cost per unit")
		fmt.Scanln(&unit)
		sl[strings.Title(item)] = shoppingList{category: index, quantity: qty, unitCost: unit}
		fmt.Println(strings.Title(item), sl[item], "has been added")

	} else {
		fmt.Println("Nothing has been added\n")
	}
}

//function to check for if entry is an existing category, if yes return category index

func ifExist(cat1 string) (index int, result bool) {

	var i int

	for i = 0; i < len(cat); i++ {
		if strings.EqualFold(cat[i], cat1) {
			index = i
			result = true
			break
		}
	}

	if i == len(cat) {
		result = false
	}
	return index, result
}

func onlyExistingCategory(cat1 string) (index int) {

	var exist bool

	for exist != true { //Only accept existing category
		index, exist = ifExist(cat1)
		if exist == true {
			break
		}
		fmt.Println("No such category exist!")
		fmt.Println("Select from existing category")
		for j := range cat {
			fmt.Println(cat[j])
		}
		fmt.Scanln(&cat1)

	}
	return
}
