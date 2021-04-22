package main

var (
	cat                 []string
	shoppingListContent = make(map[string]shoppingList)
)

type slc map[string]shoppingList

type shoppingList struct {
	category int
	quantity int
	unitCost int
}

func main() {

	cat = []string{"Household", "Food", "Drinks"}

	shoppingListContent["Cups"] = shoppingList{category: 0, quantity: 5, unitCost: 3}
	shoppingListContent["Cake"] = shoppingList{category: 1, quantity: 3, unitCost: 1}
	shoppingListContent["Sprite"] = shoppingList{category: 2, quantity: 5, unitCost: 2}
	shoppingListContent["Fork"] = shoppingList{category: 0, quantity: 4, unitCost: 3}
	shoppingListContent["Bread"] = shoppingList{category: 1, quantity: 2, unitCost: 2}
	shoppingListContent["Plates"] = shoppingList{category: 0, quantity: 4, unitCost: 3}
	shoppingListContent["Coke"] = shoppingList{category: 2, quantity: 5, unitCost: 2}

	printMenu()

}
