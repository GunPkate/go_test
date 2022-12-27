package rental

import "fmt"

type Customer struct {
	name    string
	rentals []Rental
}

func NewCustomer(name string) (rcvr Customer) {
	return Customer{
		name:    name,
		rentals: []Rental{},
	}
}

func (r Rental) Charge() float64 { //one param = one dependency
	result := 0.0
	switch r.Movie().PriceCode() {
	case REGULAR:
		result += 2
		if r.DaysRented() > 2 {
			result += float64(r.DaysRented()-2) * 1.5
		}
	case NEW_RELEASE:
		result += float64(r.DaysRented()) * 3.0
	case CHILDRENS:
		result += 1.5
		if r.DaysRented() > 3 {
			result += float64(r.DaysRented()-3) * 1.5
		}
	}
	return result
}

func getPoints(r Rental) int {
	if r.Movie().PriceCode() == NEW_RELEASE && r.DaysRented() > 1 {
		return 2
	}
	return 1
}

func (c Customer) AddRental(arg Rental) {
	c.rentals = append(c.rentals, arg)
}
func (c Customer) Name() string {
	return c.name
}

func getTotalPoint(c Customer) int {
	frequentRenterPoints := 0
	for _, r := range c.rentals {
		frequentRenterPoints += getPoints(r)
	}
	return frequentRenterPoints
}

func getTotalAmount(rentals []Rental) float64 {
	totalAmount := 0.0
	for _, r := range rentals {
		totalAmount += r.Charge()
	}
	return totalAmount
}

func (c Customer) Statement() string { //!!!!!!HTML
	totalAmount := getTotalAmount(c.rentals)
	points := getTotalPoint(c)

	result := fmt.Sprintf("Rental Record for %v\n", c.Name())
	for _, r := range c.rentals {
		title := r.Movie().Title()
		amount := r.Charge()
		result += fmt.Sprintf("\t%v\t%.1f\n", title, amount)
	}
	result += fmt.Sprintf("Amount owed is %.1f\n", totalAmount)
	result += fmt.Sprintf("You earned %v frequent renter points", points)
	return result
}
