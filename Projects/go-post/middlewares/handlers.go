package middlewares

import (
	"database/sql"
	"encoding/json"
	"example/go-post/models"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func createConnection() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// connect to postgres and check for err
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}

	// verify the db connection
	err = db.Ping()
	if err != nil {
		panic(err)

	}

	// print success message if everything works fine
	fmt.Println("Successfully connected to postgres")
	return db
}

// Create stock controller Method -> POST

func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock

	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}
	insertToDB := insertDB(stock)

	res := response{
		ID:      insertToDB,
		Message: "stock created successfully",
	}
	json.NewEncoder(w).Encode(res)

}

// Get stock controller based on the id of the stock Method -> GET
func GetStock() {}

// Get all the available stock Method -> GET
func GetAllStock() {}

// Update a stock based on passed id in query Method ->PATCH
func UpdateStock() {}

// Delete a particular stock Method -> DElETE
func DeleteStock() {}
