// Exercise 3: Inventory Management System
// Topics Covered: Go Conditions, Go Type Casting, Go Functions, Go Arrays, Go Strings,
// Go Errors
// Case Study:
// A store needs to manage its inventory of products. Build an application that includes
// the following:
// 1. Product Struct: Create a struct to represent a product with fields for ID, name,
// price (float64), and stock (int).
// 2. Add Product: Write a function to add new products to the inventory. Use type
// casting to ensure price inputs are converted to float64.
// 3. Update Stock: Implement a function to update the stock of a product. Use
// conditions to validate the input (e.g., stock cannot be negative).
// 4. Search Product: Allow users to search for products by name or ID. If a product is
// not found, return a custom error message.
// 5. Display Inventory: Use loops to display all available products in a formatted
// table.
// Bonus:
// • Add sorting functionality to display products by price or stock in ascending order.

package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

type Product struct {
	ID    string
	Name  string
	Price float64
	Stock int
}

var inventory []Product

func main() {
	for {
		fmt.Println("\nInventory Management System")
		fmt.Println("1. Add Product")
		fmt.Println("2. Update Stock")
		fmt.Println("3. Search Product")
		fmt.Println("4. Display Inventory")
		fmt.Println("5. Sort by Price")
		fmt.Println("6. Sort by Stock")
		fmt.Println("7. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}

		switch choice {
		case 1:
			addProduct()
		case 2:
			updateStock()
		case 3:
			searchProduct()
		case 4:
			displayInventory()
		case 5:
			sortByPrice()
		case 6:
			sortByStock()
		case 7:
			fmt.Println("Exiting the system. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}

func addProduct() {
	var id, name string
	var price float64
	var stock int

	fmt.Print("Enter Product ID: ")
	fmt.Scan(&id)
	fmt.Print("Enter Product Name: ")
	fmt.Scan(&name)
	fmt.Print("Enter Product Price: ")
	_, err := fmt.Scan(&price)
	if err != nil {
		fmt.Println("Invalid input for price. Please enter a valid number.")
		return
	}
	fmt.Print("Enter Product Stock: ")
	_, err = fmt.Scan(&stock)
	if err != nil || stock < 0 {
		fmt.Println("Invalid stock value. Please enter a non-negative number.")
		return
	}

	for _, product := range inventory {
		if product.ID == id {
			fmt.Println("Product ID already exists. Please use a unique ID.")
			return
		}
	}

	newProduct := Product{
		ID:    id,
		Name:  name,
		Price: price,
		Stock: stock,
	}
	inventory = append(inventory, newProduct)
	fmt.Println("Product added successfully!")
}

func updateStock() {
	var id string
	var stock int

	fmt.Print("Enter Product ID: ")
	fmt.Scan(&id)

	product, index, err := findProductByID(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Print("Enter new stock value: ")
	_, err = fmt.Scan(&stock)
	if err != nil || stock < 0 {
		fmt.Println("Invalid stock value. Please enter a non-negative number.")
		return
	}

	inventory[index].Stock = stock
	fmt.Printf("Stock updated for product '%s'. New stock: %d\n", product.Name, stock)
}

func searchProduct() {
	var search string
	fmt.Print("Enter Product Name or ID: ")
	fmt.Scan(&search)

	for _, product := range inventory {
		if product.ID == search || product.Name == search {
			fmt.Printf("Product Found: ID: %s, Name: %s, Price: %.2f, Stock: %d\n",
				product.ID, product.Name, product.Price, product.Stock)
			return
		}
	}

	fmt.Println("Error: Product not found.")
}

func displayInventory() {
	if len(inventory) == 0 {
		fmt.Println("No products in inventory.")
		return
	}

	fmt.Println("\nInventory:")
	fmt.Printf("%-10s %-20s %-10s %-10s\n", "ID", "Name", "Price", "Stock")
	fmt.Println(strings.Repeat("-", 50))
	for _, product := range inventory {
		fmt.Printf("%-10s %-20s %-10.2f %-10d\n", product.ID, product.Name, product.Price, product.Stock)
	}
}

func sortByPrice() {
	if len(inventory) == 0 {
		fmt.Println("No products to sort.")
		return
	}

	sort.Slice(inventory, func(i, j int) bool {
		return inventory[i].Price < inventory[j].Price
	})
	fmt.Println("Inventory sorted by price.")
	displayInventory()
}

func sortByStock() {
	if len(inventory) == 0 {
		fmt.Println("No products to sort.")
		return
	}

	sort.Slice(inventory, func(i, j int) bool {
		return inventory[i].Stock < inventory[j].Stock
	})
	fmt.Println("Inventory sorted by stock.")
	displayInventory()
}

func findProductByID(id string) (*Product, int, error) {
	for index, product := range inventory {
		if product.ID == id {
			return &product, index, nil
		}
	}
	return nil, -1, errors.New("product not found")
}
