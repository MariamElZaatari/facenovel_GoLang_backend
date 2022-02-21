package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type FriendRequest struct {
	Friend_request_id int    `json:"friend_request_id"`
	Requester_id      int    `json:"requester_id"`
	Receiver_id       int    `json:"receiver_id"`
	Status            string `json:"status"`
	Date_created      string `json:"date_created"`
	Date_updated      string `json:"date_updated"`
}

type FriendRequest_user struct {
	Requester_id int    `json:"requester_id"`
	First_name   string `json:"first_name"`
	Last_name    string `json:"last_name"`
	Profile_pic  string `json:"profile_pic"`
	Bio_text     string `json:"bio_text"`
}

func friendRequestCreate(c *gin.Context) {

	var newfriendRequest FriendRequest

	// Call BindJSON to Bind the Received JSON to newUser.
	if err := c.BindJSON(&newfriendRequest); err != nil {
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
	insert, err := db.Query("INSERT INTO friend_request VALUES (NULL,?,?,?,?,?)", newfriendRequest.Requester_id, newfriendRequest.Receiver_id, newfriendRequest.Status, currentTime.Format("2006-01-02 15:04:05"), currentTime.Format("2006-01-02 15:04:05"))

	// Error Inserting
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	c.IndentedJSON(http.StatusCreated, newfriendRequest)
}

func friendRequestUpdate(c *gin.Context) {

	var newfriendRequest FriendRequest

	// Call BindJSON to Bind the Received JSON to newUser.
	if err := c.BindJSON(&newfriendRequest); err != nil {
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
	insert, err := db.Query("UPDATE `friend_request` SET `status`=?`date_updated`=? WHERE requester_id=? AND receiver_id=?", newfriendRequest.Status, currentTime.Format("2006-01-02 15:04:05"), newfriendRequest.Requester_id, newfriendRequest.Receiver_id)

	// Error Updating
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	c.IndentedJSON(http.StatusCreated, newfriendRequest)
}

func friendRequestDelete(c *gin.Context) {

	var newfriendRequest FriendRequest

	// Call BindJSON to Bind the Received JSON to newUser.
	if err := c.BindJSON(&newfriendRequest); err != nil {
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
	delete, err := db.Query("DELETE FROM friend_request WHERE requester_id=? AND receiver_id=?", newfriendRequest.Requester_id, newfriendRequest.Receiver_id)

	// Error Deleting
	if err != nil {
		panic(err.Error())
	}

	defer delete.Close()

	c.IndentedJSON(http.StatusOK, newfriendRequest)
}
