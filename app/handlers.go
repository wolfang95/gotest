package app

import (
	"banking/Service"
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zipcode"`
}

type CustomerHandlers struct {
	service Service.CustomerService
}
func (ch *CustomerHandlers) getAllCustomer(w http.ResponseWriter, r *http.Request) {
	/*customers := []Customer{
		{Name: "Ashish", City: "New Delhi", Zipcode: "1112233"},
		{Name: "Rob", City: "Villa", Zipcode: "10923"},


	}
	*/

	//w.Header().Add("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(customers)

	customers, _ :=ch.service.GetAllCustomer()
	if r.Header.Get("Content-Type") == "application/xml" {

		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

	//r.Header.Get("content-Type")

}
