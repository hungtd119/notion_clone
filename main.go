package main

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Page struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {

	err := godotenv.Load()
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser, dbPass, dbHost, dbPort, dbName,
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	router.GET("/pages", func(c *gin.Context) {
		rows, err := db.Query("SELECT id, title FROM pages")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var pages []map[string]interface{}
		for rows.Next() {
			var id int
			var title string
			rows.Scan(&id, &title)
			pages = append(pages, map[string]interface{}{"id": id, "title": title})
		}

		c.JSON(http.StatusOK, pages)
	})

	router.POST("/pages", func(c *gin.Context) {
		var input struct {
			Title   string `json:"title" binding:"required"`
			Content string `json:"content"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id := uuid.New().String()
		createdAt := time.Now()

		_, err := db.Exec(
			"INSERT INTO pages (id, title, content, created_at) VALUES (?, ?, ?, ?)",
			id, input.Title, input.Content, createdAt,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Page created",
			"page": Page{
				ID:        id,
				Title:     input.Title,
				Content:   input.Content,
				CreatedAt: createdAt,
			},
		})
	})

	router.Run(":8080")
}
