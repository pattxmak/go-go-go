package repository

import "errors"

type customerRepositoryMock struct {
	customers []Customer
}

func NewCustomerRepositoryMock() customerRepositoryMock {
	customers := []Customer{
		{CustomerID: 1001, Name: "Alice Smith", DateOfBirth: "1990-05-12", City: "Bangkok", Zipcode: "10110", Status: 1},
		{CustomerID: 1002, Name: "Bob Johnson", DateOfBirth: "1985-09-30", City: "Chiang Mai", Zipcode: "50000", Status: 0},
		{CustomerID: 1003, Name: "Charlie Brown", DateOfBirth: "2000-01-15", City: "Phuket", Zipcode: "83000", Status: 1},
	}

	return customerRepositoryMock{customers: customers}
}

func (r customerRepositoryMock) GetAll() ([]Customer, error) {
	return r.customers, nil
}

func (r customerRepositoryMock) GetById(id int) (*Customer, error) {

	for _, customer := range r.customers {
		if customer.CustomerID == id {
			return &customer, nil
		}
	}
	return nil, errors.New("customer not found")
}
