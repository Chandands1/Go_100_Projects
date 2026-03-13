package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Price    float64
	Quantity int
}

type Inventory struct {
	Products map[int]Product
}

func startInventory() {
	inventory := &Inventory{
		Products: make(map[int]Product),
	}

	menuInventory(inventory)
}

func menuInventory(inv *Inventory) {

	for {
		var choice int

		fmt.Println("\n======= Inventory Management System =======")
		fmt.Println("1. Create Product")
		fmt.Println("2. View Products")
		fmt.Println("3. Delete Product")
		fmt.Println("4. Update Product")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		fmt.Scanln(&choice)

		switch choice {

		case 1:
			createProduct(inv)

		case 2:
			viewProducts(inv)

		case 3:
			deleteProduct(inv)

		case 4:
			updateProduct(inv)

		case 5:
			fmt.Println("Exiting program...")
			return

		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}

func createProduct(inv *Inventory) {

	p := Product{}

	fmt.Print("Enter Product ID: ")
	fmt.Scanln(&p.Id)

	_, exists := inv.Products[p.Id]

	if exists {
		fmt.Println("Product with this ID already exists")
		return
	}

	fmt.Print("Enter Product Name: ")
	fmt.Scanln(&p.Name)

	fmt.Print("Enter Price: ")
	fmt.Scanln(&p.Price)

	fmt.Print("Enter Quantity: ")
	fmt.Scanln(&p.Quantity)

	inv.Products[p.Id] = p

	fmt.Println("Product added successfully!")
}

func viewProducts(inv *Inventory) {

	if len(inv.Products) == 0 {
		fmt.Println("No products in inventory")
		return
	}

	fmt.Println("\n----- Product List -----")

	for id, product := range inv.Products {

		fmt.Println("Product ID:", id)
		fmt.Println("Name:", product.Name)
		fmt.Println("Price:", product.Price)
		fmt.Println("Quantity:", product.Quantity)
		fmt.Println("-----------------------")
	}
}

func updateProduct(inv *Inventory) {

	var id int

	fmt.Print("Enter Product ID to update: ")
	fmt.Scanln(&id)

	product, exists := inv.Products[id]

	if !exists {
		fmt.Println("Product not found")
		return
	}

	fmt.Print("Enter New Name: ")
	fmt.Scanln(&product.Name)

	fmt.Print("Enter New Price: ")
	fmt.Scanln(&product.Price)

	fmt.Print("Enter New Quantity: ")
	fmt.Scanln(&product.Quantity)

	inv.Products[id] = product

	fmt.Println("Product updated successfully!")
}

func deleteProduct(inv *Inventory) {

	var id int

	fmt.Print("Enter Product ID to delete: ")
	fmt.Scanln(&id)

	_, exists := inv.Products[id]

	if !exists {
		fmt.Println("Product not found")
		return
	}

	delete(inv.Products, id)

	fmt.Println("Product deleted successfully!")
}

func main() {
	startInventory()
}