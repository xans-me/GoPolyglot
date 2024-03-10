# GoPolyglot

GoPolyglot is a Go (Golang) library designed to facilitate message translations and dynamic translations within Go applications. It provides an efficient way to manage multilingual support, particularly useful in web applications.
## Features

- **Translation of response messages** based on Response Code Identifiers (RC).
- **Support for dynamic translations** tied to specific endpoints, allowing for real-time translation adjustments.
- **Flexible and customizable configuration** for mapping database columns and crafting SQL queries tailored to your schema.
- **Framework-agnostic middleware** capable of integrating with popular Go web frameworks such as Gin and Mux, making it versatile for various web applications.

## Installation

To use GoPolyglot in your Go project, you need to install it using `go get`:

```bash
go install github.com/xans-me/GoPolyglot
```

## Usage

Below is an example demonstrating how to integrate GoPolyglot with a Gin router. The setup for Mux follows a similar pattern, thanks to the agnostic nature of the provided middleware.
```go
package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xans-me/GoPolyglot"
	"github.com/xans-me/GoPolyglot/middleware"
	"net/http"
)

func main() {
	router := gin.Default()
	db, err := sql.Open("mysql", "your-database-connection-string")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	config := GoPolyglot.Config{
		// Configuration for your translation schema
	}

	translator := GoPolyglot.NewTranslator(db, config)
	router.Use(middleware.GinWrapper(middleware.AgnosticTranslationMiddleware(translator)))

	router.GET("/example", func(c *gin.Context) {
		// Your handler code
	})

	router.Run(":8080")
}

```
### For Mux Users
If you are using Mux, the middleware can be directly applied without wrapping:

```go
router := mux.NewRouter()
router.Use(middleware.AgnosticTranslationMiddleware(translator))

```

## Configuration

Define your translation schema and queries in the Config struct:


```go
config := GoPolyglot.Config{
ColumnConfig: GoPolyglot.ColumnConfig{
TableName:         "your_translation_table",
RCColumn:          "response_code",
TitleColumn:       "title",
DescriptionColumn: "description",
},
// Additional configuration options
}

translator := GoPolyglot.NewTranslator(db, config)

```


## Middleware
The library includes middleware compatible with both Gin and Mux for translating response data based on the request's parameters and headers.

## License
This library is released under the MIT License. See the LICENSE file for details.

## Contributing
Contributions are welcome! If you find any issues or have suggestions for improvements, please create an issue or a pull request.

## Author
Teuku Mulia Ichsan

## Acknowledgments
Special thanks to the Go community.

