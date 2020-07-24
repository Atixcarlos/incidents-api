package controllers

import (
	"encoding/json"
	"incidents-api/models/employee"
	"incidents-api/utils"
	"net/http"
)

// EmployeesHandler is a funtion to handle diferent methods
func EmployeesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Serve the resource.
		ListEmployees(w, r)
	case http.MethodPost:
		// Create a new record.
		SaveEmployee(w, r)
	default:
		// Give an error message.
		w.WriteHeader(http.StatusNotImplemented)
	}
}

// ListEmployees is a function to connect model.
func ListEmployees(w http.ResponseWriter, r *http.Request) {

	// Get data from context
	myCustomerID := r.Context().Value("customerID")
	employees, err := employee.List(myCustomerID.(int))
	if err != nil {
		utils.LogError(r, err)
	}
	data, err := json.Marshal(employees)
	if err != nil {
		utils.InternalServerErrorWriter(w)
	} else {
		utils.WriteJSON(w, http.StatusOK, data)
	}

}

// SaveEmployee is a function to connect model.
func SaveEmployee(w http.ResponseWriter, r *http.Request) {

	// Get data from context
	myCustomerID := r.Context().Value("customerID")
	myEmployee := employee.New()

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&myEmployee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set custom attributes
	myEmployee.CustomerID = myCustomerID.(int)

	err = myEmployee.Save()
	if err != nil {
		utils.LogError(r, err)
	} else {
		utils.WriteJSON(w, http.StatusCreated, []byte(`{"success": true, "message": "Success, Employee yas added correctly."}`))
	}

}
