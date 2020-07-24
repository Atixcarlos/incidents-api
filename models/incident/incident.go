package incident

import (
	"incidents-api/database"
	"time"
)

// Incident defines a incident
type Incident struct {
	ID         int       `json:"id"`
	Type       string    `json:"type"`
	StartDate  time.Time `json:"startDate"`
	EndDate    time.Time `json:"endDate"`
	Note       string    `json:"note"`
	EmployeeId int       `json:"employeeId"`
}

// List loads all employees.
func List() ([]Incident, error) {

	myData := make([]Incident, 0)

	rows, err := database.DBCon.Query("SELECT * FROM incidents.employees")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		i := new(Incident)
		err = rows.Scan(&i.ID, &i.Type, &i.StartDate, &i.EndDate, &i.Note, &i.EmployeeId)
		if err != nil {
			return nil, err
		}
		myData = append(myData, *i)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	rows.Close()

	return myData, nil
}

// Load loads an incident from dbase.
func Load(id int) (*Incident, error) {

	i := new(Incident)

	err := database.DBCon.QueryRow("SELECT * FROM incidents.incidents WHERE id = $1", id).
		Scan(&i.ID, &i.Type, &i.StartDate, &i.EndDate, &i.Note, &i.EmployeeId)
	if err != nil {
		return nil, err
	}

	return i, nil
}

// Save saves an incident to the dbase.
func (i *Incident) Save() error {

	err := database.DBCon.QueryRow("INSERT into incidents.incidents(type, start_date, end_date, note, employee_id) VALUES($1, $2, $3, $4, $5)", i.Type, i.StartDate, i.EndDate, i.Note, i.EmployeeId).Scan()

	if err != nil {
		return err
	}

	return nil
}
