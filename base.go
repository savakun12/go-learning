package main

import (
    "fmt"
    "errors"
)

// Struct untuk merepresentasikan sebuah buku
type Book struct {
    Title  string
    Author string
    Pages  int
}

// Method untuk struct Book
func (b Book) Summary() string {
    return fmt.Sprintf("%s by %s, %d pages", b.Title, b.Author, b.Pages)
}

// Fungsi untuk membuat buku baru dengan validasi
func NewBook(title, author string, pages int) (*Book, error) {
    if title == "" {
        return nil, errors.New("title cannot be empty")
    }
    if author == "" {
        return nil, errors.New("author cannot be empty")
    }
    if pages <= 0 {
        return nil, errors.New("pages must be a positive number")
    }
    return &Book{Title: title, Author: author, Pages: pages}, nil
}

func main() {
    // Membuat beberapa buku
    book1, err := NewBook("The Go Programming Language", "Alan A. A. Donovan", 380)
    if err != nil {
        fmt.Println("Error creating book1:", err)
        return
    }

    book2, err := NewBook("", "Unknown", 100)
    if err != nil {
        fmt.Println("Error creating book2:", err)
    }

    // Menggunakan method Summary
    fmt.Println(book1.Summary())

    // Slice of books
    books := []Book{*book1, {"Go in Action", "William Kennedy", 265}}

    // Iterasi melalui slice
    for _, book := range books {
        fmt.Println(book.Summary())
    }
}
