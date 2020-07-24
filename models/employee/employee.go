package employee

import (
	"fmt"
	"incidents-api/database"
)

// Employee defines a employee
type Employee struct {
	ID         int    `json:"id"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	CustomerID int    `json:"customerID"`
}

// New returns a new *Employee variable.
func New() *Employee {
	return new(Employee)
}

// List loads all employees.
func List(customerID int) ([]Employee, error) {

	myData := make([]Employee, 0)
	rows, err := database.DBCon.Query("SELECT id, first_name, last_name, email, customer_id FROM incidents.employees where customer_id=$1", customerID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		e := New()
		err = rows.Scan(&e.ID, &e.FirstName, &e.LastName, &e.Email, &e.CustomerID)
		if err != nil {
			return nil, err
		}
		myData = append(myData, *e)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	rows.Close()

	return myData, nil
}

// Load loads an employee from dbase.
func Load(id int) (*Employee, error) {

	e := New()

	err := database.DBCon.QueryRow("SELECT id, first_name, last_name, email, customer_id FROM incidents.employees WHERE id = $1", id).
		Scan(&e.ID, &e.FirstName, &e.LastName, &e.Email, &e.CustomerID)
	if err != nil {
		return nil, err
	}

	return e, nil
}

// Save saves an employee to the dbase.
func (e *Employee) Save() error {

	err := database.DBCon.QueryRow("INSERT into incidents.employees(first_name, last_name, email, customer_id) VALUES($1, $2, $3, $4)", e.FirstName, e.LastName, e.Email, e.CustomerID).Scan()

	if err != nil {
		return err
	}

	return nil
}
