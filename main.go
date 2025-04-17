package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

// Book struct
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var (
	books  = make(map[int]Book)
	nextID = 1
	mu     sync.Mutex
)

func main() {
	http.HandleFunc("/books", booksHandler)     // GET all or POST
	http.HandleFunc("/books/", bookByIDHandler) // GET, PUT, DELETE by ID

	fmt.Println("🚀 Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// GET /books or POST /books
func booksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getAllBooks(w, r)
	case "POST":
		createBook(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// GET /books/{id} or PUT or DELETE
func bookByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/books/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "GET":
		getBook(w, r, id)
	case "PUT":
		updateBook(w, r, id)
	case "DELETE":
		deleteBook(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var bookList []Book
	for _, b := range books {
		bookList = append(bookList, b)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookList)
}

func getBook(w http.ResponseWriter, r *http.Request, id int) {
	mu.Lock()
	defer mu.Unlock()

	book, exists := books[id]
	if !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil || book.Title == "" || book.Author == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	mu.Lock()
	book.ID = nextID
	nextID++
	books[book.ID] = book
	mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request, id int) {
	var updated Book
	err := json.NewDecoder(r.Body).Decode(&updated)
	if err != nil || updated.Title == "" || updated.Author == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	_, exists := books[id]
	if !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	updated.ID = id
	books[id] = updated
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updated)
}

func deleteBook(w http.ResponseWriter, r *http.Request, id int) {
	mu.Lock()
	defer mu.Unlock()

	_, exists := books[id]
	if !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	delete(books, id)
	w.WriteHeader(http.StatusNoContent)
}
