package customer

import (
	"strings"
)

type Customer struct {
	Index            string `json:"index"`
	CustomerId       string `json:"customerId"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	Company          string `json:"company"`
	City             string `json:"city"`
	Country          string `json:"country"`
	Phone1           string `json:"phone1"`
	Phone2           string `json:"phone2"`
	Email            string `json:"email"`
	SubscriptionDate string `json:"subscriptionDate"`
	Website          string `json:"website"`
}

func CreateCustomer(str string) Customer {
	s := strings.Split(str, ",")
	return Customer{
		Index:            s[0],
		CustomerId:       s[1],
		FirstName:        s[2],
		LastName:         s[3],
		Company:          s[4],
		City:             s[5],
		Country:          s[6],
		Phone1:           s[7],
		Phone2:           s[8],
		SubscriptionDate: s[9],
		Website:          s[10],
	}
}
