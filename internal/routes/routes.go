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
	mux.HandleFunc("PUT /customer/{email}", updateCustomer)
	mux.HandleFunc("DELETE /customer/{email}", deleteCustomer)

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

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	email := r.PathValue("email")
	decodeBody := json.NewDecoder(r.Body)
	customer := new(Customer)
	err := decodeBody.Decode(&customer)

	if err != nil {
		fmt.Println("err", err)
	}
	_, err = database.DBCon.Exec("UPDATE CUSTOMERS SET FIRST_NAME = :1, LAST_NAME = :2, SEX = :3 WHERE EMAIL = :4", 
	customer.FirstName, customer.LastName, customer.Sex, email)

	 response := []byte(fmt.Sprintf("Successfully updated customer  %s", email))
	w.Write(response)


	if err != nil {
		panic(err)
	}
}


func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	email := r.PathValue("email")

	_, err := database.DBCon.Exec("DELETE FROM CUSTOMERS WHERE EMAIL = :1", email)

	if err != nil {
		panic(err)
	}

	 response := []byte(fmt.Sprintf("Successfully deleted customer  %s", email))
	w.Write(response)


}