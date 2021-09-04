package main

import (
	"fmt"
	"time"
)

// Global variable
var test = "blah"

func main() {
	customer := getData(1)
	customer2 := getData(2)

	fmt.Println(customer)
	fmt.Println(customer2)
	fmt.Println(getData(3))
	fmt.Println(test)
	fmt.Println(getCustomers())
	fmt.Println(len(getCustomers()))
	fmt.Println(getSliceOfCustomers())

	customers := GetCustomers()

	for _, c := range customers {
		fmt.Println(c)
	}

	// infinite loop
	for {
		fmt.Println("Infinite Loop 1")
		time.Sleep(time.Second)
		break
	}
}

// func getData(customerID int) (customer string) {
// 	firstname := "Andy"
// 	var lastname = "Godish"

// 	fullname := firstname + " " + lastname

// 	return fullname
// }

func getData(customerID int) (customer string) {
	if customerID == 1 {
		return "Andy Godish"
	} else if customerID == 2 {
		return "Bob Smith"
	} else {
		return " "
	}
}

func getCustomers() (customers [2]string) {
	customer := "New Customer"

	customers[0] = customer
	customers[1] = "Someone else"

	return customers
}

func getSliceOfCustomers() (customers []string) {
	// customers = []string{} // this is now a slice
	// or
	customers = []string{"item 1", "item 2"}

	customers = append(customers, "Andy G")
	customers = append(customers, "Another Customer")

	// for x := 0; x < len(customers); x++ { // increment int variables
	// 	fmt.Println(customers[x])
	// }

	for x, _ := range customers {
		fmt.Println((customers[x]))
	}

	return customers
}
