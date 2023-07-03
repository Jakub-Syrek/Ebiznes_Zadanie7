package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/rs/cors"
)

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Payment struct {
	ID         string  `json:"id"`
	Amount     float64 `json:"amount"`
	CardNumber string  `json:"cardNumber"`
	CardExpiry string  `json:"cardExpiry"`
	CardCVV    string  `json:"cardCvv"`
}

var products = []Product{
	{ID: "1", Name: "Product 1", Price: 10.00},
	{ID: "2", Name: "Product 2", Price: 20.00},
	{ID: "3", Name: "Product 3", Price: 30.00},
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handlePayments(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var payment Payment
		err = json.Unmarshal(bodyBytes, &payment)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Printf("Received payment: %+v\n", payment)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(payment)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/products", getProducts)
	mux.HandleFunc("/api/payments", handlePayments)

	fmt.Println("Starting server")
	handler := cors.Default().Handler(mux)
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		fmt.Println("Fatal Error")
		log.Fatal(err)
	}
}
