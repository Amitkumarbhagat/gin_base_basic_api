package main

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Amit Bhagat", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "Sumit bhagat", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Ankit bhagat", Quantity: 6},
}
