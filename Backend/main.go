package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// album represents data about a record album.
type Attractions struct {
	Id         string `db:"id" json:"id"`
	Name       string `db:"name" json:"name"`
	Detail     string `db:"detail" json:"detail"`
	Coverimage string `db:"coverimage" json:"coverimage"`
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "root:Mmcc0741@tcp(localhost:3306)/gomysql")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/attractions", getAttactions)

	router.Run("localhost:8080")
}

// getattractions responds with the list of all attractions as JSON.
func getAttactions(c *gin.Context) {
	var attractions []Attractions
	rows, err := db.Query("select id, name ,detail, coverimage from attractions ")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var a Attractions
		err := rows.Scan(&a.Id, &a.Name, &a.Detail, &a.Coverimage)
		if err != nil {
			log.Fatal(err)
		}
		attractions = append(attractions, a)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, attractions)
}
