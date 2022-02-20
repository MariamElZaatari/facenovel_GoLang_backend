package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Education struct {
	Education_id int    `json:"education_id"`
	User_id      int    `json:"user_id"`
	School_name  string `json:"school_name"`
	Date_from    string `json:"date_from"`
	Date_to      string `json:"date_to"`
	Date_created string `json:"date_created"`
	Date_updated string `json:"date_updated"`
}

func educationCreate(c *gin.Context) {

	var newEducation Education

	// Call BindJSON to bind the received JSON to
	// newEducation.
	if err := c.BindJSON(&newEducation); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "bad Input")
		return
	}

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	//username:password@tcp(127.0.0.1:3306)/DBname
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/facebookdb")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()
	// perform a db.Query insert
	currentTime := time.Now()
	insert, err := db.Query("INSERT INTO education VALUES (NULL,?,?,?,?,?,?)", newEducation.User_id, newEducation.School_name, newEducation.Date_from, newEducation.Date_to, currentTime.Format("2006-01-02 15:04:05"), currentTime.Format("2006-01-02 15:04:05"))

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()

	c.IndentedJSON(http.StatusCreated, newEducation)
}

func educationUpdate(c *gin.Context) {

	var newEducation Education

	// Call BindJSON to bind the received JSON to
	// newEducation.
	if err := c.BindJSON(&newEducation); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "bad Input")
		return
	}

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	//username:password@tcp(127.0.0.1:3306)/DBname
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/facebookdb")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()
	// perform a db.Query insert
	currentTime := time.Now()
	insert, err := db.Query("UPDATE `education` SET `school_name`=?,`date_from`=?,`date_to`=?, `date_updated`=? WHERE `education_id`=?", newEducation.School_name, newEducation.Date_from, newEducation.Date_to, currentTime.Format("2006-01-02 15:04:05"), newEducation.Education_id)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()

	c.IndentedJSON(http.StatusCreated, newEducation)
}

func educationDelete(c *gin.Context) {

	id := c.Param("id")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	//username:password@tcp(127.0.0.1:3306)/DBname
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/facebookdb")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()
	// perform a db.Query delete
	delete, err := db.Query("DELETE FROM education WHERE education_id=?", id)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer delete.Close()

	c.IndentedJSON(http.StatusOK, id)
}

func educationReadByUserID(c *gin.Context) {

	id := c.Param("id")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	//username:password@tcp(127.0.0.1:3306)/DBname
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/facebookdb")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT * FROM education WHERE user_id=? ORDER BY date_from DESC, date_to DESC, school_name", id)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var educations_user = []Education{}

	for results.Next() {
		var education_user Education
		// for each row, scan the result into our tag composite object
		err = results.Scan(&education_user.Education_id, &education_user.User_id, &education_user.School_name, &education_user.Date_from, &education_user.Date_to, &education_user.Date_created, &education_user.Date_updated)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		educations_user = append(educations_user, education_user)
	}

	c.IndentedJSON(http.StatusOK, educations_user)
}
