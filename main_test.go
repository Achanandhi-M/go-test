package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"strings"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test generated using Keploy
func TestCreateBook_ValidInput_003(t *testing.T) {
	// Arrange
	books = make(map[int]Book) // Reset books map
	nextID = 1
	newBook := Book{Title: "New Book", Author: "New Author"}
	body, err := json.Marshal(newBook)
	require.NoError(t, err)

	req, err := http.NewRequest("POST", "/books", strings.NewReader(string(body)))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	// Act
	createBook(rr, req)

	// Assert
	assert.Equal(t, http.StatusCreated, rr.Code)
	var createdBook Book
	err = json.Unmarshal(rr.Body.Bytes(), &createdBook)
	require.NoError(t, err)
	assert.Equal(t, newBook.Title, createdBook.Title)
	assert.Equal(t, newBook.Author, createdBook.Author)
	assert.Equal(t, 1, createdBook.ID)
	assert.Equal(t, createdBook, books[1])
}

// Test generated using Keploy

func TestUpdateBook_BookNotFound_009(t *testing.T) {
	// Arrange
	books = make(map[int]Book) // Reset books map
	updatedBook := Book{Title: "Updated Title", Author: "Updated Author"}
	body, err := json.Marshal(updatedBook)
	require.NoError(t, err)

	req, err := http.NewRequest("PUT", "/books/1", strings.NewReader(string(body)))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	// Act
	updateBook(rr, req, 1)

	// Assert
	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Equal(t, "Book not found\n", rr.Body.String())
}

// Test generated using Keploy

func TestGetBook_BookExists_005(t *testing.T) {
	// Arrange
	books = map[int]Book{
		1: {ID: 1, Title: "Book One", Author: "Author One"},
	}
	req, err := http.NewRequest("GET", "/books/1", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()

	// Act
	getBook(rr, req, 1)

	// Assert
	assert.Equal(t, http.StatusOK, rr.Code)
	var book Book
	err = json.Unmarshal(rr.Body.Bytes(), &book)
	require.NoError(t, err)
	assert.Equal(t, books[1], book)
}

// Test generated using Keploy

// Test generated using Keploy

func TestDeleteBook_BookExists_010(t *testing.T) {
	// Arrange
	books = map[int]Book{
		1: {ID: 1, Title: "Book One", Author: "Author One"},
	}
	req, err := http.NewRequest("DELETE", "/books/1", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()

	// Act
	deleteBook(rr, req, 1)

	// Assert
	assert.Equal(t, http.StatusNoContent, rr.Code)
	_, exists := books[1]
	assert.False(t, exists)
}

// Test generated using Keploy

func TestGetBook_BookNotFound_006(t *testing.T) {
	// Arrange
	books = make(map[int]Book) // Reset books map
	req, err := http.NewRequest("GET", "/books/1", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()

	// Act
	getBook(rr, req, 1)

	// Assert
	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Equal(t, "Book not found\n", rr.Body.String())
}

// Test generated using Keploy

func TestCreateBook_InvalidInput_007(t *testing.T) {
	// Arrange
	invalidBody := `{"title": "", "author": ""}`
	req, err := http.NewRequest("POST", "/books", strings.NewReader(invalidBody))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	// Act
	createBook(rr, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, "Invalid input\n", rr.Body.String())
}

// Test generated using Keploy

func TestDeleteBook_BookNotFound_011(t *testing.T) {
	// Arrange
	books = make(map[int]Book) // Reset books map
	req, err := http.NewRequest("DELETE", "/books/1", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()

	// Act
	deleteBook(rr, req, 1)

	// Assert
	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Equal(t, "Book not found\n", rr.Body.String())
}

func TestGetAllBooks_MultipleBooks_002(t *testing.T) {
	// Arrange
	books = map[int]Book{
		1: {ID: 1, Title: "Book One", Author: "Author One"},
		2: {ID: 2, Title: "Book Two", Author: "Author Two"},
	}
	req, err := http.NewRequest("GET", "/books", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()

	// Act
	getAllBooks(rr, req)

	// Assert
	assert.Equal(t, http.StatusOK, rr.Code)
	var bookList []Book
	err = json.Unmarshal(rr.Body.Bytes(), &bookList)
	require.NoError(t, err)
	assert.Len(t, bookList, 2)
	assert.Equal(t, books[1], bookList[0])
	assert.Equal(t, books[2], bookList[1])
}

// Test generated using Keploy
