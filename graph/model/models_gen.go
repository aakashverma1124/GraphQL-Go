// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Mutation struct {
}

type NewProduct struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
	Description int    `json:"description"`
}

type Product struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
	Description string `json:"description"`
}

type Query struct {
}
