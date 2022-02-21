package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type POST struct {
	Post_id      int    `json:"post_id"`
	User_id      int    `json:"user_id"`
	Text         string `json:"text"`
	Date_created string `json:"date_created"`
}
type POST_USER struct {
	Post_id      int    `json:"post_id"`
	First_name   string `json:"first_name"`
	Last_name    string `json:"last_name"`
	Profile_pic  string `json:"profile_pic"`
	Text         string `json:"text"`
	Date_created string `json:"date_created"`
	NumOfLikes   int    `json:"numOfLikes"`
}

func postRead(c *gin.Context) {

	// Database Connection.
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/facebookdb")

	// Error Opening The Connection
	if err != nil {
		panic(err.Error())
	}

	// Defer Close Connection Till After Query Executing
	defer db.Close()

	// Execute The Query
	results, err := db.Query("SELECT * FROM post")
	if err != nil {
		panic(err.Error())
	}

	var posts = []POST{}

	for results.Next() {
		var post POST
		// Scan Row in Result Into The Tag Composite Object
		err = results.Scan(&post.Post_id, &post.User_id, &post.Text, &post.Date_created)
		if err != nil {
			panic(err.Error())
		}
		posts = append(posts, post)
	}

	c.IndentedJSON(http.StatusOK, posts)
}

func postCreate(c *gin.Context) {

	var newPost POST

	// Call BindJSON to Bind the Received JSON to newUser.
	if err := c.BindJSON(&newPost); err != nil {
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
	insert, err := db.Query("INSERT INTO post VALUES ( Null,?,?,? )", newPost.User_id, newPost.Text, currentTime.Format("2006-01-02 15:04:05"))

	// Error Inserting
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	c.IndentedJSON(http.StatusCreated, newPost)
}

func postDelete(c *gin.Context) {

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
	delete, err := db.Query("DELETE FROM post WHERE post_id= ?", id)

	// Error Deleting
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer delete.Close()

	c.IndentedJSON(http.StatusOK, id)
}

func postReadByID(c *gin.Context) {

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
	results, err := db.Query("SELECT p.post_id, u.first_name, u.last_name, u.profile_pic, p.text, p.date_created, Count(l.likes_id) as numOfLikes FROM post as p LEFT JOIN likes as l ON p.post_id=l.post_id JOIN user_info as u ON p.user_id=u.user_id WHERE u.user_id=? GROUP BY p.post_id ORDER BY date_created DESC", id)

	// Error Selecting
	if err != nil {
		panic(err.Error())
	}

	var posts_user = []POST_USER{}

	for results.Next() {
		var post_user POST_USER
		// Scan Row in Result Into The Tag Composite Object
		err = results.Scan(&post_user.Post_id, &post_user.First_name, &post_user.Last_name, &post_user.Profile_pic, &post_user.Text, &post_user.Date_created, &post_user.NumOfLikes)
		if err != nil {
			panic(err.Error())
		}
		posts_user = append(posts_user, post_user)
	}

	c.IndentedJSON(http.StatusOK, posts_user)
}
