package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
)

type City struct {
	ID          int    `json:"ID,omitempty" db:"ID"`
	Name        string `json:"name,omitempty" db:"Name"`
	CountryCode string `json:"countryCode,omitempty"  db:"CountryCode"`
	District    string `json:"district,omitempty"  db:"District"`
	Population  int    `json:"population,omitempty"  db:"Population"`
}

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True&loc=Asia%%2FTokyo&charset=utf8mb4",
		os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOSTNAME"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("conntected")

	newCity := City{
		Name:        "Oookayama",
		CountryCode: "JPN",
		District:    "Tokyo-to",
		Population:  5000,
	}
	result, err := db.Exec("INSERT INTO city (Name, CountryCode, District, Population) VALUES (?,?,?,?)", newCity.Name, newCity.CountryCode, newCity.District, newCity.Population)
	if err != nil {
		log.Fatal(err)
	}

	id, _ :=  result.LastInsertId()

	fmt.Printf("%sのIDは%d\n", newCity.Name, id)
}
