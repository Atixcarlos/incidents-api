package controllers

import (
	"encoding/json"
	"incidents-api/models/employee"
	"incidents-api/utils"
	"net/http"
	"strconv"
)

// EmployeeHandler is a funtion to handle diferent methods
func EmployeeHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/employees/"):]
	myID, _ := strconv.Atoi(id)

	switch r.Method {
	case http.MethodGet:
		// Serve the resource.
		LoadEmployee(myID, w, r)
	case http.MethodPut:
		// Update an existing record.

	case http.MethodDelete:
		// Remove the record.

	default:
		// Give an error message.
		w.WriteHeader(http.StatusNotImplemented)
	}
}

// LoadEmployee is a function to connect model.
func LoadEmployee(myID int, w http.ResponseWriter, r *http.Request) {
	// Get data from context
	employees, err := employee.Load(myID)

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

