package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type auth_user struct {
	User_id         string `json:"user_id"`
	First_name      string `json:"first_name"`
	Last_name       string `json:"last_name"`
	Email           string `json:"email"`
	Username        string `json:"username"`
	Phone           string `json:"phone"`
	Gender          string `json:"gender"`
	Dob             string `json:"dob"`
	Password        string `json:"password"`
	Password_repeat string `json:"password_repeat"`
	Date_created    string `json:"date_created"`
	Date_updated    string `json:"date_updated"`
}

func signup(c *gin.Context) {

	var newUser auth_user

	// Call BindJSON to bind the received JSON to
	// newUser.
	if err := c.BindJSON(&newUser); err != nil {
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
	insert, err := db.Query("INSERT INTO user(user_id,username, password, date_created, date_updated) VALUES (NULL,?,?,?,?)", newUser.Username, newUser.Password, currentTime.Format("2006-01-02 15:04:05"), currentTime.Format("2006-01-02 15:04:05"))

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()

	c.IndentedJSON(http.StatusCreated, newUser)
}

func login(c *gin.Context) {

	var user auth_user

	// Call BindJSON to bind the received JSON to
	// newUser.
	if err := c.BindJSON(&user); err != nil {
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

	var loggedUser auth_user

	// Execute the query
	err = db.QueryRow("SELECT * FROM user WHERE username = ? AND password=?", user.Username, user.Password).Scan(&loggedUser.User_id, &loggedUser.Username, &loggedUser.Password, &loggedUser.Date_created, &loggedUser.Date_updated)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	c.IndentedJSON(http.StatusOK, loggedUser)
}
