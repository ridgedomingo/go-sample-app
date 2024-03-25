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

	mux.HandleFunc("GET /customers", getAllCustomers)
	mux.HandleFunc("POST /customer", createCustomer)

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

func createCustomer(w http.ResponseWriter, r *http.Request) {
	decodeBody := json.NewDecoder(r.Body)
	customer := new(Customer)
	err := decodeBody.Decode(&customer)

	if err != nil {
		fmt.Println("err", err)
	}

	_, err = database.DBCon.Exec("INSERT INTO CUSTOMERS VALUES (:1, :2, :3, :4)", customer.Email, customer.FirstName, customer.LastName, customer.Sex)

	 response := []byte(fmt.Sprintf("Successfully created customer %s %s", customer.FirstName,  customer.LastName ))
	w.Write(response)


	if err != nil {
		panic(err)
	}
}