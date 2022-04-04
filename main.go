package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type DayToShipCount map[int]int

var ShipCount DayToShipCount

func main() {
	ShipCount = make(DayToShipCount)
	ShipCount[1] = 100
	ShipCount[2] = 50

	router := http.NewServeMux()
	router.HandleFunc("/ship", GetShipCount)
	router.HandleFunc("/ship/create", SetShipCount)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", router)
	log.Fatal(err)
}

func GetShipCount(w http.ResponseWriter, r *http.Request) {
	day, _ := strconv.Atoi(r.URL.Query().Get("day"))
	if day <= 0 {
		http.Error(w, "Day cannot be less than zero", 400)
		return
	}

	count, ok := ShipCount[day]
	if ok == false {
		http.Error(w, "Day cannot found in map", 400)
		return
	}

	fmt.Fprintf(w, "Ship Count=%d \n", count)
}

// curl -X POST "http://localhost:4000/ship/create?day=4&count=20"
func SetShipCount(w http.ResponseWriter, r *http.Request) {
	day, _ := strconv.Atoi(r.URL.Query().Get("day"))
	if day <= 0 {
		http.Error(w, "Count cannot be zero or less than zero", 400)
		return
	}

	count, _ := strconv.Atoi(r.URL.Query().Get("count"))
	if count <= 0 {
		http.Error(w, "Count cannot be zero or less than zero", 400)
		return
	}

	ShipCount[day] = count
	w.WriteHeader(200)
}
