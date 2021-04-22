package main

import "fmt"

func (sl slc) printCurrentData() {

	fmt.Println("Print Current Data.")
	for k, v := range sl {
		if k != "" {
			fmt.Println(k, "-", v)
		} else {
			fmt.Println("No data found!")
		}
	}

}
