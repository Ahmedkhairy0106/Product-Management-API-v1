package main

import (
	"net/http"
	"slices"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"

)

type Product struct {
	ID string
	Title string
	Description string
	Price float64
	Quantity int
}

var products = []Product{
	{ID: "1", Title: "Laptop", Description: "Dell Laptop", Price: 45000, Quantity: 10},
	{ID: "2", Title: "Mobile", Description: "Samsung Mobile", Price: 25000, Quantity: 20},
	{ID: "3", Title: "Tablet", Description: "Lenovo Tablet", Price: 15000, Quantity: 30},
	{ID: "4", Title: "Headphones", Description: "Sony Headphones", Price: 1000, Quantity: 50},
	{ID: "5", Title: "Mouse", Description: "Logitech Mouse", Price: 500, Quantity: 70},
	{ID: "6", Title: "Keyboard", Description: "Logitech Keyboard", Price: 700, Quantity: 100},
	{ID: "7", Title: "Monitor", Description: "Asus Monitor", Price: 12000, Quantity: 80},
	{ID: "8", Title: "Printer", Description: "HP Printer", Price: 8000, Quantity: 40},
	{ID: "9", Title: "Scanner", Description: "Canon Scanner", Price: 5000, Quantity: 60},
	{ID: "10", Title: "Projector", Description: "Epson Projector", Price: 25000, Quantity: 10},
	{ID: "11", Title: "Mousepad", Description: "Logitech Mousepad", Price: 300, Quantity: 150},
	{ID: "12", Title: "Speaker", Description: "Sony Speaker", Price: 1500, Quantity: 250},
	{ID: "13", Title: "Pendrive", Description: "Sandisk Pendrive", Price: 400, Quantity: 200},
	{ID: "14", Title: "Harddisk", Description: "Seagate Harddisk", Price: 3000, Quantity: 100},
	{ID: "15", Title: "Router", Description: "TP-Link Router", Price: 1500, Quantity: 50},
	{ID: "16", Title: "Switch", Description: "TP-Link Switch", Price: 500, Quantity: 100},
	{ID: "17", Title: "Server", Description: "Dell Server", Price: 12000, Quantity: 20},
	{ID: "18", Title: "UPS", Description: "APC UPS", Price: 5000, Quantity: 30},
	{ID: "19", Title: "Webcam", Description: "Logitech Webcam", Price: 2000, Quantity: 40},
	{ID: "20", Title: "Cooler", Description: "Cooler Master Cooler", Price: 1000, Quantity: 50},
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	if len(products) == 0 {
		render.JSON(w, r, "No products found")
		return
	}else{
		render.JSON(w, r, products)
	}
}

func getProductByID(w http.ResponseWriter, r *http.Request){
	productID := chi.URLParam(r, "id")
	for _, product := range products {
		if productID == product.ID {
			render.JSON(w, r, product)
			return
		}
	}
	render.JSON(w, r, "Product not found")
}

func addproduct(w http.ResponseWriter, r *http.Request){
	var newProduct Product
	booL := false
	// render.Decode is a function that reads the incoming data from the request (r) and converts it from JSON into a struct in Go.
	err := render.Decode(r, &newProduct)
	if err != nil {
		render.JSON(w, r, "Invalid request")
		return
	}

	if newProduct.ID == "" || newProduct.Title == "" || newProduct.Description == "" || newProduct.Price == 0 || newProduct.Quantity == 0 {
		render.JSON(w, r, "Invalid request")
		return
	}

	for _, product := range products {
		if newProduct.ID == product.ID {
			render.JSON(w, r, "the Product id is already exists")
			booL = true
			return
		}
	}
	if !booL {
		products = append(products, newProduct)
	}
	render.JSON(w, r, products)
}

func updateproduct(w http.ResponseWriter, r *http.Request) {
    productID := chi.URLParam(r, "id")

    var updatedProduct Product
    err := render.Decode(r, &updatedProduct)
    if err != nil {
        render.JSON(w, r, "Invalid request")
        return
    }
	// Ensure that the updated product ID does not already exist in the list 
	// while allowing the update if the ID remains unchanged.
    for _, product := range products {
        if product.ID == updatedProduct.ID && product.ID != productID {
            render.JSON(w, r, "Product ID already exists")
            return
        }
    }

    for i, product := range products {
        if product.ID == productID {
            if updatedProduct.Title != "" {
                product.Title = updatedProduct.Title
            }
            if updatedProduct.Description != "" {
                product.Description = updatedProduct.Description
            }
            if updatedProduct.Price != 0 {
                product.Price = updatedProduct.Price
            }
            if updatedProduct.Quantity != 0 {
                product.Quantity = updatedProduct.Quantity
            }
            if updatedProduct.ID != "" {
                product.ID = updatedProduct.ID
            }

            products[i] = product
            render.JSON(w, r, products)
            return
        }
    }

    render.JSON(w, r, "Product not found")
}


func deleteproduct(w http.ResponseWriter, r *http.Request){
		productID := chi.URLParam(r, "id")
		for i, product := range products {
			// We can use the slices package to delete an element from a slices
			// The slices package is a custom package that we have created to delete an element from a slice
			// what we have here first one which is the slice we want to delete from, the second one is the start index of the element we want to delete and the
			// third one is the end index of the element we want to delete
			if productID == product.ID {
				products = slices.Delete(products, i, i+1)
				render.JSON(w, r, products)
				return
			}
		}
		render.JSON(w, r, "Product not found")
}

func main() {
	ch := chi.NewRouter()
	ch.Use(middleware.Logger)

	// chi.Router is the value returned by chi.NewRouter()
	ch.Route("/api", func (api chi.Router){

		api.Get("/products", getProduct)
		api.Get("/products/{id}", getProductByID)

		api.Post("/products", addproduct)

		api.Put("/products/{id}", updateproduct)
		api.Delete("/products/{id}", deleteproduct)
	})

	http.ListenAndServe(":8080", ch)
}