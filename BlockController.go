package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Block struct {
	Block_id        int    `json:"block_id"`
	User_id         int    `json:"user_id"`
	Blocked_user_id int    `json:"blocked_user_id"`
	Date_created    string `json:"date_created"`
}

type Block_USER struct {
	Blocked_user_id int    `json:"blocked_user_id"`
	First_name      string `json:"first_name"`
	Last_name       string `json:"last_name"`
	Profile_pic     string `json:"profile_pic"`
	Bio_text        string `json:"bio_text"`
}

func blockCreate(c *gin.Context) {

	var newBlock Block

	// Call BindJSON to Bind the Received JSON to newUser.
	if err := c.BindJSON(&newBlock); err != nil {
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
	insert, err := db.Query("INSERT INTO block VALUES ( Null,?,?,? )", newBlock.User_id, newBlock.Blocked_user_id, currentTime.Format("2006-01-02 15:04:05"))

	// Error Inserting
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	c.IndentedJSON(http.StatusCreated, newBlock)
}

func blockDelete(c *gin.Context) {

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
	delete, err := db.Query("DELETE FROM block WHERE block_id= ?", id)

	// Error Deleting
	if err != nil {
		panic(err.Error())
	}

	defer delete.Close()

	c.IndentedJSON(http.StatusOK, id)
}

func blockReadByUserID(c *gin.Context) {

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
	results, err := db.Query("SELECT b.blocked_user_id, u.first_name, u.last_name, u.profile_pic, u.bio_text FROM `block` as b JOIN user_info as u ON b.blocked_user_id=u.user_id WHERE b.user_id = ? ORDER BY b.date_created", id)
	
	// Error Selecting
	if err != nil {
		panic(err.Error())
	}

	var blocks_user = []Block_USER{}

	for results.Next() {
		var block_user Block_USER
		// Scan Row in Result Into The Tag Composite Object
		err = results.Scan(&block_user.Blocked_user_id, &block_user.First_name, &block_user.Last_name, &block_user.Profile_pic, &block_user.Bio_text)
		
		if err != nil {
			panic(err.Error())
		}

		blocks_user = append(blocks_user, block_user)
	}

	c.IndentedJSON(http.StatusOK, blocks_user)
}