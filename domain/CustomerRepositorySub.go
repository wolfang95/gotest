package domain
type CustomerRepositorySub struct {
	customers []Customer
}

func (s CustomerRepositorySub) FindAll() ([]Customer,error){
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositorySub{
	customers := []Customer{
		{Id:"1001",Name:"Wolfang", City:"New Delhi", Zipcode:"110011",DateofBirth:"2000-01-01"},
		{Id:"1002",Name:"Rob", City:"New De", Zipcode:"110011",DateofBirth:"2000-01-01"},
	}
	return CustomerRepositorySub{customers: customers}
}