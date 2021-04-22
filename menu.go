package main

import "fmt"

func printMenu() {

	choice := 0

	for choice != 10 {

		fmt.Println("\nShopping List Application")
		fmt.Println("=========================")
		fmt.Println("1. View entire shopping list")
		fmt.Println("2. Generate Shopping List Report")
		fmt.Println("3. Add Items")
		fmt.Println("4. Modify Items")
		fmt.Println("5. Delete Item")
		fmt.Println("6. Print Current Data Fields")
		fmt.Println("7. Add New Category Name")
		fmt.Println("8. Modify/Delete Category")
		fmt.Println("9. Store/retrieve shopping list")
		fmt.Println("10. Stop this program")
		fmt.Println("Select your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1: //View entire shopping list
			slc(shoppingListContent).printShoppingList()
		case 2: //Generate Shopping List Report
			generateReport()
		case 3: //Add Items
			slc(shoppingListContent).addItem()
		case 4: //Modify Items
			slc(shoppingListContent).modifyItem()
		case 5: //Delete Items
			slc(shoppingListContent).deleteItem()
		case 6: //Print Current Data Fields
			slc(shoppingListContent).printCurrentData()
		case 7: //Add New Category Name
			addCategory()
		case 8: //Modify/Delete Category
			modifyDeleteCategory()
		case 9: //Store/Retrieve shopping list
			storingRetrieveList()
		case 10: //exiting the application
			fmt.Println("Exiting the shopping list application")
			break
		default:
			fmt.Println("Please select 1 to 10.")
		}
	}

}
