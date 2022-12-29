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

func RegularCharge(r Rental) float64 {
	result := 2.0
	if r.DaysRented() > 2 {
		result += float64(r.DaysRented()-2) * 1.5
	}
	return result
}

func NewReleaseCharge(r Rental) float64 {
	return float64(r.DaysRented()) * 3.0
}

func ChildrenCharge(daysRented int) float64 {
	result := 1.5
	if daysRented > 3 {
		result += float64(daysRented-3) * 1.5
	}
	return result
}

func (r Rental) Charge() float64 { //one param = one dependency
	switch r.Movie().PriceCode() {
	case REGULAR:
		return r.Movie().Charger.Charge(r.daysRented)
	case NEW_RELEASE:
		return r.Movie().Charger.Charge(r.daysRented)
	case CHILDRENS:
		return r.Movie().Charger.Charge(r.daysRented)
	case 0:
		return r.Movie().Charger.Charge(r.daysRented)
	}
	return 0
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

func getTotalPoint(rentals []Rental) int {
	points := 0
	for _, r := range rentals {
		points += getPoints(r)
	}
	return points
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
	points := getTotalPoint(c.rentals)

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
