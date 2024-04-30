package middlewares

import (
	"database/sql"
	"encoding/json"
	"example/go-post/models"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"

	"github.com/joho/godotenv"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

// func createConnection() *sql.DB {
// 	err := godotenv.Load(".env")

// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	// connect to postgres and check for err
// 	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
// 	if err != nil {
// 		panic(err)
// 	}

// 	// verify the db connection
// 	err = db.Ping()
// 	if err != nil {
// 		panic(err)

// 	}

// 	// print success message if everything works fine
// 	fmt.Println("Successfully connected to postgres")
// 	return db
// }

// DB init -> single persistent connection through out program lifecycle

var db *sql.DB

func init() {
	// Load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatalf("Error connecting to the DB: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging the DB :%v", err)
	}

	fmt.Println("DB connection successful")

}

// close the DB once program exits / closes
func closeDB() {
	if db != nil {
		db.Close()
	}
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
func GetStock(w http.ResponseWriter, r *http.Request) {
	// extract variables of current requested url path
	params := mux.Vars(r)

	// get the id field from the param
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string to int.%v", err)
	}

	// get the stock associated with the id
	stock, err := getStock(int64(id))

	if err != nil {
		log.Fatalf("Failed to get stock %v", err)
	}

	json.NewEncoder(w).Encode(stock)
}

// Get all the available stock Method -> GET
func GetAllStock(w http.ResponseWriter, r *http.Request) {

	// get all the stocks
	stocks, err := getAllStocks()

	if err != nil {
		log.Fatalf("Failed to fetch all the stocks %v", err)
	}

	json.NewEncoder(w).Encode(stocks)

}

// Update a stock based on passed id in query Method ->PATCH
func UpdateStock(w http.ResponseWriter, r *http.Request) {

	// get the id from requested url path with mux.Vars
	params := mux.Vars(r)

	// extract the id from the
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to process conversion of the requested id %v", err)
	}

	// create an instance of the stocks model
	var stock models.Stock

	// decode the json request to stock
	err = json.NewDecoder(r.Body).Decode(&stock)

	if err != nil {
		log.Fatalf("unable the decode the request body from update stocks %v", err)

	}

	// if ok! call the updateStock function to update the stock

	updatedRows := updateSingleStock(int64(id), stock)

	msg := fmt.Sprintf("Stock updated successfully. Total rows/record affected %v", updatedRows)

	// send the response object
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)

}

// Delete a particular stock Method -> DElETE
func DeleteStock(w http.ResponseWriter, r *http.Request) {
	// get the stockid from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id in string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	// call the deleteStock, convert the int to int64
	deletedRows := deleteStock(int64(id))

	// format the message string
	msg := fmt.Sprintf("Stock updated successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)

}

// function handler to query with postgress

func insertDB(stock models.Stock) int64 {

	sqlQuery := `insert into stocks(name,price,company) values($1,$2,$3) returning stockid`

	var id int64

	err := db.QueryRow(sqlQuery, stock.Name, stock.Price, stock.Company).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)
	return id
}

// function to get stock -> Query
func getStock(id int64) (models.Stock, error) {
	// instance of the model to query
	var stock models.Stock
	// sql query to get the stock

	sqlQuery := `select * from stocks where stockid=$1`

	// execute the stament
	row := db.QueryRow(sqlQuery, id)
	err := row.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows returned")
		return stock, nil
	case nil:
		return stock, nil
	default:
		log.Fatalf("Unable to scan the row  %v", err)
	}

	return stock, err

}

// func to getAllStock

func getAllStocks() ([]models.Stock, error) {
	var stocks []models.Stock
	// query statement to execute
	sqlQuery := `select * from stocks`
	// passing this query to database/sql package to fetch the result
	rows, err := db.Query(sqlQuery)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var stock models.Stock
		err := rows.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)
		if err != nil {
			return nil, err
		}
		stocks = append(stocks, stock)
	}
	return stocks, nil

}

// Update the stock ->

func updateSingleStock(id int64, stock models.Stock) int64 {
	sqlQuery := `UPDATE stocks SET name=$2, price=$3, company=$4 WHERE stockid=$1`

	// execute the query
	res, err := db.Exec(sqlQuery, id, stock.Name, stock.Price, stock.Company)

	if err != nil {
		log.Fatalf("Unable to execute the update query : %v", err)
	}

	updatedRows, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows: %v", err)
	}
	fmt.Printf("Total rows/record affected  %v\n", updatedRows)
	return updatedRows

}

// query function to delete the element

func deleteStock(id int64) int64 {
	sqlQuery := `delete from stock where stockid =$1`

	res, err := db.Exec(sqlQuery, id)
	if err != nil {
		log.Fatalf("Unable to process the delete query: %v", err)
	}
	deletedRows, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected/deleted rows: %v", err)
	}
	fmt.Printf("Total rows/record affected %v\n", deletedRows)
	return deletedRows
}
