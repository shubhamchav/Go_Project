// handlers/todo_handler_test.go
package handlers

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestGetAllTodos(t *testing.T) {
    req, err := http.NewRequest("GET", "/todos", nil)
    if err != nil {
        t.Fatal(err)
    }

    // Create a response recorder to record the response
    rr := httptest.NewRecorder()

    // Create an HTTP handler from the handler function
    handler := http.HandlerFunc(todoHandler.GetAllTodos)

    // Serve the HTTP request to the recorder
    handler.ServeHTTP(rr, req)

    // Check the status code
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // Check the response body
    expected := `[{"id":1,"title":"Example Todo","description":"Example Description","completed":false}]`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

// Write similar test functions for other handlers (GetTodoByID, CreateTodo, UpdateTodo, DeleteTodo)
