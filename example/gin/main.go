package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver
	"github.com/xans-me/GoPolyglot"
	"github.com/xans-me/GoPolyglot/middleware"
)

func main() {
	// Create a Gin router with the default middleware.
	router := gin.Default()

	// Open a MySQL database connection.
	db, err := sql.Open("mysql", "")
	if err != nil {
		panic(err) // Handle error by panicking
	}
	defer db.Close() // Ensure the database connection is closed when the function returns

	// Configure translation settings.
	config := GoPolyglot.Config{
		ColumnConfig: GoPolyglot.ColumnConfig{
			TableName:         "",            // Name of the table containing translations
			RCColumn:          "rc",          // Column name for response codes
			TitleColumn:       "title",       // Column name for titles
			DescriptionColumn: "description", // Column name for descriptions
		},
	}

	// Create a translator instance with the configured database and translation settings.
	translator := GoPolyglot.NewTranslator(db, config)

	// Use the translation middleware with the router.
	router.Use(GoPolyglot.GinWrapper(middleware.AgnosticTranslationMiddleware(translator)))

	// Define a route that returns a response with translation.
	router.GET("/rc", func(c *gin.Context) {
		response := map[string]interface{}{
			"rc":          "RZP1",
			"trxType":     "Login",
			"trxFeature":  "Login",
			"title":       "This is Title",
			"description": "This is description",
		}

		c.JSON(http.StatusUnauthorized, response) // Send a JSON response with status code 401
	})

	// Start the Gin server on port 8080.
	router.Run(":8080")
}
