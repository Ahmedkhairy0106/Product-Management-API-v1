# 📦 Product Management API

![Product Management API Banner](https://img.shields.io/badge/Product%20Management%20API-v1.0-blueviolet?style=for-the-badge&logo=go)  
![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat-square&logo=go)  
![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)

---

This is a Go-based Product Management API that provides functionality for managing products in an inventory system. It supports the following operations:

- Get all products
- Get product by ID
- Add a new product
- Update product details
- Delete a product

## API Endpoints

### 1. Get All Products

`GET /api/products`

This endpoint returns a list of all products in the system.

**Response:**
```json
[
  {
    "ID": "1",
    "Title": "Laptop",
    "Description": "Dell Laptop",
    "Price": 45000,
    "Quantity": 10
  },
  ...
]
```

### 2. Get Product By ID

`GET /api/products/{id}`

This endpoint returns a product with the given ID.

**Parameters:**
- `id`: The ID of the product to retrieve.

**Response:**
```json
{
  "ID": "1",
  "Title": "Laptop",
  "Description": "Dell Laptop",
  "Price": 45000,
  "Quantity": 10
}
```

### 3. Add New Product

`POST /api/products`

This endpoint adds a new product to the system.

**Request Body:**
```json
{
  "ID": "16",
  "Title": "Switch",
  "Description": "TP-Link Switch",
  "Price": 500,
  "Quantity": 100
}
```

**Response:**
The updated list of products.

**Error Responses:**
- `400`: Invalid request if any required field is missing or the product ID already exists.

### 4. Update Product Details

`PUT /api/products/{id}`

This endpoint updates the details of an existing product.

**Parameters:**
- `id`: The ID of the product to update.

**Request Body:**
```json
{
  "Title": "New Title",
  "Description": "Updated Description",
  "Price": 600,
  "Quantity": 80
}
```

**Response:**
The updated list of products.

**Error Responses:**
- `404`: Product not found if the product with the given ID does not exist.

### 5. Delete Product

`DELETE /api/products/{id}`

This endpoint deletes the product with the given ID from the system.

**Parameters:**
- `id`: The ID of the product to delete.

**Response:**
The updated list of products.

**Error Responses:**
- `404`: Product not found if the product with the given ID does not exist.

## Conclusion

This API provides basic functionality for managing products in an inventory system. It supports adding, updating, deleting, and retrieving product data.
