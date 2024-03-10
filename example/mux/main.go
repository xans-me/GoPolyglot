package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"         // Import the MySQL driver
	"github.com/gorilla/mux"                   // Use Gorilla Mux for routing
	"github.com/xans-me/GoPolyglot"            // Import the GoPolyglot package for translation
	"github.com/xans-me/GoPolyglot/middleware" // Import the middleware package from GoPolyglot
	"net/http"
)

func main() {
	// Create a router with Gorilla Mux
	router := mux.NewRouter()

	// Open a MySQL database connection
	db, err := sql.Open("mysql", "")
	if err != nil {
		panic(err) // Handle any errors that occur during database connection setup
	}
	defer db.Close() // Ensure the database connection is closed when the main function returns

	// Configure translation settings
	config := GoPolyglot.Config{
		ColumnConfig: GoPolyglot.ColumnConfig{
			TableName:         "",            // Name of the table containing translations
			RCColumn:          "rc",          // Column name for response codes
			TitleColumn:       "title",       // Column name for titles
			DescriptionColumn: "description", // Column name for descriptions
		},
	}

	// Create a translator instance with the configured database and translation settings
	translator := GoPolyglot.NewTranslator(db, config)

	// Apply the translation middleware to the Mux router
	router.Use(middleware.AgnosticTranslationMiddleware(translator))

	// Define a route that returns a response with translation
	router.HandleFunc("/rc", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]interface{}{
			"rc":          "RZP1",
			"trxType":     "Login",
			"trxFeature":  "Login",
			"title":       "This is Title",
			"description": "This is description",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
	})
	
	// Start the Mux server on port 8080
	http.ListenAndServe(":8080", router)
}
