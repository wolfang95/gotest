package app

import (
	"banking/Service"
	"banking/domain"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func Start() {

	//mux := http.NewServeMux()
	router := mux.NewRouter()

	//writing
	//ch := CustomerHandlers{ Service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	ch := CustomerHandlers{ Service.NewCustomerService(domain.NewCustomerRepositoryDB())}

	//define routes
	router.HandleFunc("/customers", ch.getAllCustomer).Methods(http.MethodGet)
	//router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustumer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post Request was recived")
}

func getCustumer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w,vars["customer_id"])
}
