package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ridgedomingo/go-sample-app/internal/database"
)

type Customer struct {
    Email string
    FirstName string
	LastName string
	Sex string
}


func NewRouter () http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/customers", getAllCustomers)

	return mux
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	var customers []*Customer
	rows, err := database.DBCon.Query("SELECT * FROM CUSTOMERS")
	if err != nil {
		fmt.Println("Error while getting customers", err)
	}

	for rows.Next() {
		customer := new(Customer)
		err := rows.Scan(&customer.Email, &customer.FirstName, &customer.LastName, &customer.Sex)
		if err != nil {
			fmt.Println(err)
		}
		customers = append(customers, customer)
	}

	if err := json.NewEncoder(w).Encode(customers); err != nil {
		fmt.Println(err)
	}
}