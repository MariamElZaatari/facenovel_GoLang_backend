package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	User_info_id int    `json:"user_info_id"`
	User_id      int    `json:"user_id"`
	First_name   string `json:"first_name"`
	Last_name    string `json:"last_name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Gender       string `json:"gender"`
	Dob          string `json:"dob"`
	Profile_pic  string `json:"profile_pic"`
	Bio_text     string `json:"bio_text"`
	Current_city string `json:"current_city"`
	Date_created string `json:"date_created"`
	Date_updated string `json:"date_updated"`
}

func userCreate(c *gin.Context) {

	var newUser User

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
	insert, err := db.Query("INSERT INTO `user_info` VALUES (NULL,?,?,?,?,?,?,?,?,?,?,?,?)", newUser.User_id, newUser.First_name, newUser.Last_name, newUser.Email, newUser.Phone, newUser.Gender, newUser.Dob, newUser.Profile_pic, newUser.Bio_text, newUser.Current_city, currentTime.Format("2006-01-02 15:04:05"), currentTime.Format("2006-01-02 15:04:05"))

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()

	c.IndentedJSON(http.StatusCreated, newUser)
}

func userUpdate(c *gin.Context) {

	var newUser User

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
	insert, err := db.Query("UPDATE `user_info` SET `first_name`=?,`last_name`=?,`email`=?,`phone`=?,`gender`=?,`dob`=?,`profile_pic`=?,`bio_text`=?,`current_city`=?,`date_updated`=? WHERE `user_id`=?", newUser.First_name, newUser.Last_name, newUser.Email, newUser.Phone, newUser.Gender, newUser.Dob, newUser.Profile_pic, newUser.Bio_text, newUser.Current_city, currentTime.Format("2006-01-02 15:04:05"), newUser.User_id)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()

	c.IndentedJSON(http.StatusCreated, newUser)
}

func userDelete(c *gin.Context) {

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
	delete, err := db.Query("DELETE FROM user_info WHERE user_id=?", id)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer delete.Close()

	c.IndentedJSON(http.StatusOK, id)
}

func userRead(c *gin.Context) {

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

	var user User

	// Execute the query
	err = db.QueryRow("SELECT * FROM user_info WHERE user_id=?", id).Scan(&user.User_info_id, &user.User_id, &user.First_name, &user.Last_name, &user.Email, &user.Phone, &user.Gender, &user.Dob, &user.Profile_pic, &user.Bio_text, &user.Current_city, &user.Date_created, &user.Date_updated)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	c.IndentedJSON(http.StatusOK, user)
}
