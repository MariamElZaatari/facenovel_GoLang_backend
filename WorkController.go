package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Work struct {
	Work_id      int    `json:"work_id"`
	User_id      int    `json:"user_id"`
	Company_name string `json:"company_name"`
	Date_from    string `json:"date_from"`
	Date_to      string `json:"date_to"`
	Date_created string `json:"date_created"`
	Date_updated string `json:"date_updated"`
}

func workCreate(c *gin.Context) {

	var newWork Work

	// Call BindJSON to bind the received JSON to
	// newWork.
	if err := c.BindJSON(&newWork); err != nil {
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
	insert, err := db.Query("INSERT INTO work VALUES (NULL,?,?,?,?,?,?)", newWork.User_id, newWork.Company_name, newWork.Date_from, newWork.Date_to, currentTime.Format("2006-01-02 15:04:05"), currentTime.Format("2006-01-02 15:04:05"))

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()

	c.IndentedJSON(http.StatusCreated, newWork)
}

func workUpdate(c *gin.Context) {

	var newWork Work

	// Call BindJSON to bind the received JSON to
	// newWork.
	if err := c.BindJSON(&newWork); err != nil {
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
	insert, err := db.Query("UPDATE `work` SET `company_name`=?,`date_from`=?,`date_to`=?, `date_updated`=? WHERE `work_id`=?", newWork.Company_name, newWork.Date_from, newWork.Date_to, currentTime.Format("2006-01-02 15:04:05"), newWork.Work_id)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()

	c.IndentedJSON(http.StatusCreated, newWork)
}

func workDelete(c *gin.Context) {

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
	delete, err := db.Query("DELETE FROM work WHERE work_id=?", id)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer delete.Close()

	c.IndentedJSON(http.StatusOK, id)
}

func workReadByUserID(c *gin.Context) {

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
	results, err := db.Query("SELECT * FROM work WHERE user_id=? ORDER BY date_from DESC, date_to DESC, company_name", id)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var works_user = []Work{}

	for results.Next() {
		var work_user Work
		// for each row, scan the result into our tag composite object
		err = results.Scan(&work_user.Work_id, &work_user.User_id, &work_user.Company_name, &work_user.Date_from, &work_user.Date_to, &work_user.Date_created, &work_user.Date_updated)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		works_user = append(works_user, work_user)
	}

	c.IndentedJSON(http.StatusOK, works_user)
}
