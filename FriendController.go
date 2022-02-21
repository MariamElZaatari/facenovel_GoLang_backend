package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Friend struct {
	Friend_id    int    `json:"friend_id"`
	User_one_id  int    `json:"user_one_id"`
	User_two_id  int    `json:"user_two_id"`
	Date_created string `json:"date_created"`
}

func friendCreate(c *gin.Context) {

	var newFriend Friend

	// Call BindJSON to Bind the Received JSON to newUser.
	if err := c.BindJSON(&newFriend); err != nil {
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
	insert, err := db.Query("INSERT INTO friend VALUES ( Null,?,?,? )", newFriend.User_one_id, newFriend.User_two_id, currentTime.Format("2006-01-02 15:04:05"))

	// Error Inserting
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	c.IndentedJSON(http.StatusCreated, newFriend)
}

func friendDelete(c *gin.Context) {

	var friend Friend

	// Call BindJSON to Bind the Received JSON to newUser.
	if err := c.BindJSON(&friend); err != nil {
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
	delete, err := db.Query("DELETE FROM friend WHERE user_one_id=? AND user_two_id=?", friend.User_one_id, friend.User_two_id)

	// Error Deleting
	if err != nil {
		panic(err.Error())
	}

	defer delete.Close()

	c.IndentedJSON(http.StatusOK, friend)
}

func friendReadByUserID(c *gin.Context) {

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
	results, err := db.Query("SELECT f.user_two_id, u.first_name, u.last_name, u.profile_pic, u.bio_text FROM friend as f JOIN user_info as u ON f.user_two_id=u.user_id WHERE user_one_id = ? UNION SELECT f.user_one_id, u.first_name, u.last_name, u.profile_pic, u.bio_text FROM friend as f JOIN user_info as u ON f.user_one_id=u.user_id WHERE user_two_id = ? ORDER BY first_name, last_name", id)

	// Error Selecting
	if err != nil {
		panic(err.Error())
	}

	var friends_user = []Friend{}

	for results.Next() {
		var friend_user Friend
		// Scan Row in Result Into The Tag Composite Object
		err = results.Scan(&friend_user.Friend_id, &friend_user.User_one_id, &friend_user.User_one_id)
		if err != nil {
			panic(err.Error())
		}
		friends_user = append(friends_user, friend_user)
	}

	c.IndentedJSON(http.StatusOK, friends_user)
}
