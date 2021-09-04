package main

type Customer struct {
	FirstName string
	LastName  string
	FullName  string
}

func GetCustomers() (customers []Customer) {
	andy := Customer{FirstName: "Andy", LastName: "Andy"}

	customers = append(customers,
		andy,
		Customer{FirstName: "Some Guy", LastName: "Last"},
	)

	return customers
}
