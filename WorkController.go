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

	// Call BindJSON to Bind the Received JSON to newUser.
	if err := c.BindJSON(&newWork); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "bad Input")
		return
	}

	// Database Connection.
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/facebookdb")

	// Error Opening The Connection
	if err != nil {
		panic(err.Error())
	}

	// Defer Close Connection Till After Query Executing
	defer db.Close()

	// Execute The Query
	currentTime := time.Now()
	insert, err := db.Query("INSERT INTO work VALUES (NULL,?,?,?,?,?,?)", newWork.User_id, newWork.Company_name, newWork.Date_from, newWork.Date_to, currentTime.Format("2006-01-02 15:04:05"), currentTime.Format("2006-01-02 15:04:05"))

	// Error Inserting
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	c.IndentedJSON(http.StatusCreated, newWork)
}

func workUpdate(c *gin.Context) {

	var newWork Work

	// Call BindJSON to Bind the Received JSON to newUser.
	if err := c.BindJSON(&newWork); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "bad Input")
		return
	}

	// Database Connection.
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/facebookdb")

	// Error Opening The Connection
	if err != nil {
		panic(err.Error())
	}

	// Defer Close Connection Till After Query Executing
	defer db.Close()

	// Execute The Query
	currentTime := time.Now()
	insert, err := db.Query("UPDATE `work` SET `company_name`=?,`date_from`=?,`date_to`=?, `date_updated`=? WHERE `work_id`=?", newWork.Company_name, newWork.Date_from, newWork.Date_to, currentTime.Format("2006-01-02 15:04:05"), newWork.Work_id)

	// Error Updating
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	c.IndentedJSON(http.StatusCreated, newWork)
}

func workDelete(c *gin.Context) {

	id := c.Param("id")

	// Database Connection.
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/facebookdb")

	// Error Opening The Connection
	if err != nil {
		panic(err.Error())
	}

	// Defer Close Connection Till After Query Executing
	defer db.Close()

	// Execute The Query
	delete, err := db.Query("DELETE FROM work WHERE work_id=?", id)

	// Error Deleting
	if err != nil {
		panic(err.Error())
	}

	defer delete.Close()

	c.IndentedJSON(http.StatusOK, id)
}

func workReadByUserID(c *gin.Context) {

	id := c.Param("id")

	// Database Connection.
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/facebookdb")

	// Error Opening The Connection
	if err != nil {
		panic(err.Error())
	}

	// Defer Close Connection Till After Query Executing
	defer db.Close()

	// Execute The Query
	results, err := db.Query("SELECT * FROM work WHERE user_id=? ORDER BY date_from DESC, date_to DESC, company_name", id)

	if err != nil {
		panic(err.Error())
	}

	var works_user = []Work{}

	for results.Next() {
		var work_user Work
		// Scan Row in Result Into The Tag Composite Object
		err = results.Scan(&work_user.Work_id, &work_user.User_id, &work_user.Company_name, &work_user.Date_from, &work_user.Date_to, &work_user.Date_created, &work_user.Date_updated)
		if err != nil {
			panic(err.Error())
		}
		works_user = append(works_user, work_user)
	}

	c.IndentedJSON(http.StatusOK, works_user)
}
