package main

import (
	"database/sql"
	"incidents-api/controllers"
	"incidents-api/database"
	"incidents-api/middlewares"
	"incidents-api/utils"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {

	// load the app config
	utils.LoadConfig("./config.json")

	// Database abstraction.
	var err error
	database.DBCon, err = sql.Open("postgres", utils.Config.DBConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer database.DBCon.Close()

	// Create server logs and maintain log file rotation.
	utils.InitLogs(utils.Config.LogDirectory)

	commonMiddleware := []Middleware{
		middlewares.LogMiddleware,
		middlewares.BasicAuthentication,
	}

	// Add handle funcs.
	http.HandleFunc("/", http.HandlerFunc(controllers.Hello))
	http.HandleFunc("/employees", MultipleMiddleware(controllers.EmployeesHandler, commonMiddleware...))
	http.HandleFunc("/employees/", MultipleMiddleware(controllers.EmployeeHandler, commonMiddleware...))
	// Run the web server.
	utils.ErrorLog.Fatal(http.ListenAndServe(utils.Config.NetworkAddr, nil))
}
