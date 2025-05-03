package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		// fmt.Println("Failed3")
		log.Fatal("Error loading .env file")
	}
	host = "localhost"
	databaseName = os.Getenv("POSTGRES_DB")
	username = os.Getenv("POSTGRES_USER")
	password = os.Getenv("POSTGRES_PASSWORD")
}

var (
	host         string
	port         = 5432
	databaseName string
	username     string
	password     string
)

var db *sql.DB

type Product struct {
	ID    int
	Name  string
	Price int
}

func main() {
	fmt.Println(os.Getenv("POSTGRES_DB"))
	fmt.Println(databaseName)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, username, password, databaseName)

	sdb, err := sql.Open("postgres", psqlInfo)
	fmt.Println(psqlInfo)
	if err != nil {
		// fmt.Println("failed1")
		log.Fatal(err)
	}

	db = sdb

	err = db.Ping()

	if err != nil {
		// fmt.Println("failed2")
		log.Fatal(err)
	}

	// Connection Database Successfull
	fmt.Println("Connection Database Successfull")

	err = createProduct(&Product{Name: "Go product3", Price: 700})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Create Successfull!")
}

func createProduct(product *Product) error {
	//INSERT INTO public.products(id, name, price) VALUES (?, ?, ?);

	_, err := db.Exec(
		"INSERT INTO public.products(name, price) VALUES ($1, $2);",
		product.Name,
		product.Price,
	)

	return err
}
