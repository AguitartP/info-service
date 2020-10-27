package main

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"log"
//	"fmt"
//	"io/ioutil"
)

var db *sql.DB
var err error

type Cars struct {
	ID int
	Brand string 
	Model string 
	Horse int
}

/*
func getCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var cars []Cars

	result, err := db.Query("SELECT id, brand, model, horse_power FROM cars")
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {
		var car Cars
		err := result.Scan(&car.ID, &car.Brand, &car.Model, &car.Horse)
		if err != nil {
			panic(err.Error())
		}
		cars = append(cars, car)
	}

	json.NewEncoder(w).Encode(cars)
}*/

func getCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	result, err := db.Query("SELECT id, brand, model, horse_power FROM cars WHERE id = ?", params["id"])
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var car Cars

	for result.Next() {
		err := result.Scan(&car.ID, &car.Brand, &car.Model, &car.Horse)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(car)
}


func main() {
	db, err = sql.Open("mysql", "postter:12345@tcp(my-mysql:3306)/test")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	router := mux.NewRouter().StrictSlash(true)

//	router.HandleFunc("/cars", getCars).Methods("GET")
	router.HandleFunc("/cars/{id}", getCar).Methods("GET")
//	router.HandleFunc("/cars", createCar).Methods("POST")

	log.Fatal(http.ListenAndServe(":8081", router))

}