package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

type App struct {
	db *sql.DB
}

func (app *App) migrate() {
	db := app.db
	down := `DROP TABLE IF EXISTS todos, activities;`
	todoUp := `CREATE TABLE IF NOT EXISTS todos (
		todo_id INT NOT NULL AUTO_INCREMENT,
		activity_group_id INT NOT NULL,
		title VARCHAR(255) NOT NULL,
		is_active BOOLEAN NOT NULL,
		priority VARCHAR(255) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		PRIMARY KEY (todo_id)
	);
`
	activityUp := `CREATE TABLE IF NOT EXISTS activities (
		activity_id INT NOT NULL AUTO_INCREMENT,
		title VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		PRIMARY KEY (activity_id)
);`

	_, err := app.db.Exec(down)
	if err != nil {
		log.Println("Failed to execute query:", err)
	}
	_, err = db.Exec(todoUp)
	if err != nil {
		log.Println("Failed to execute query:", err)
		return
	}

	_, err = db.Exec(activityUp)
	if err != nil {
		log.Println("Failed to execute query:", err)
		return
	}

	log.Println("Table created successfully")
}

func (app *App) Start(e *echo.Echo) {
	app.migrate()
	// app.echo.Logger.Fatal(app.echo.Start(":3030"))
	e.Logger.Fatal(e.Start(":8090"))

}

func LoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		url := req.URL.Path
		method := req.Method
		// Read the request body
		body, err := io.ReadAll(req.Body)
		if err != nil {
			log.Printf("%s %s (error reading body): %v", method, url, err)
		} else {
			log.Printf("%s %s %s", method, url, string(body))
		}

		// Reset the request body to the original state
		req.Body = io.NopCloser(bytes.NewBuffer(body))

		return next(c)
	}
}

func main() {

	mysqlHost := os.Getenv("MYSQL_HOST")
	if mysqlHost == "" {
		mysqlHost = "192.168.0.104"
	}

	mysqlUser := os.Getenv("MYSQL_USER")
	if mysqlUser == "" {
		mysqlUser = "root"
	}

	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	if mysqlPassword == "" {
		mysqlPassword = "faris"
	}

	mysqlDBName := os.Getenv("MYSQL_DBNAME")
	if mysqlDBName == "" {
		mysqlDBName = "skyshi"
	}

	db_cred := fmt.Sprintf("%s:%s@tcp(%s)/%s", mysqlUser, mysqlPassword, mysqlHost, mysqlDBName)
	// db_cred := "root:faris@tcp(192.168.0.104:3306)/skyshi"

	conn, err := sql.Open("mysql", db_cred)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	err = conn.Ping()
	if err != nil {
		panic(err.Error())
	}

	log.Println("Successfully connected to the database!")

	e := echo.New()

	app := App{
		db: conn,
		// echo: e,
	}
	app.getRoutes(e)
	// e.Use(LoggerMiddleware)
	app.Start(e)

}
