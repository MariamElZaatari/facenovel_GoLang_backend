package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Friend struct {
	Friend_id int    `json:"friend_id"`
	User_one_id  int    `json:"user_one_id"`
	User_two_id  int `json:"user_two_id"`
	Date_created  string `json:"date_created"`
}

func friendCreate(c *gin.Context) {

	var newFriend Friend

	// Call BindJSON to bind the received JSON to
	// newFriend.
	if err := c.BindJSON(&newFriend); err != nil {
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
	currentTime := time.Now()
	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO friend VALUES ( Null,?,?,? )", newFriend.User_one_id, newFriend.User_two_id, currentTime.Format("2006-01-02 15:04:05"))

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()

	c.IndentedJSON(http.StatusCreated, newFriend)
}

func friendDelete(c *gin.Context) {

	var friend Friend

	// Call BindJSON to bind the received JSON to
	// newFriend.
	if err := c.BindJSON(&friend); err != nil {
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
	// perform a db.Query delete
	delete, err := db.Query("DELETE FROM friend WHERE user_one_id=? AND user_two_id=?", friend.User_one_id, friend.User_two_id)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer delete.Close()

	c.IndentedJSON(http.StatusOK, friend)
}

func friendReadByUserID(c *gin.Context) {

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
	results, err := db.Query("SELECT f.user_two_id, u.first_name, u.last_name, u.profile_pic, u.bio_text FROM friend as f JOIN user_info as u ON f.user_two_id=u.user_id WHERE user_one_id = ? UNION SELECT f.user_one_id, u.first_name, u.last_name, u.profile_pic, u.bio_text FROM friend as f JOIN user_info as u ON f.user_one_id=u.user_id WHERE user_two_id = ? ORDER BY first_name, last_name", id)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var friends_user = []Friend{}

	for results.Next() {
		var friend_user Friend
		// for each row, scan the result into our tag composite object
		err = results.Scan(&friend_user.Friend_id, &friend_user.User_one_id, &friend_user.User_one_id)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		friends_user = append(friends_user, friend_user)
	}
	// and then print out the tag's Name attribute
	// log.Printf(posts)

	c.IndentedJSON(http.StatusOK, friends_user)
}

