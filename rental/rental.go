package rental

type Rental struct {
	movie      Movie
	daysRented int
}

func NewRental(movie Movie, daysRented int) (rcvr Rental) {
	rcvr = Rental{
		movie:      movie,
		daysRented: daysRented,
	}
	return
}

func (rcvr Rental) DaysRented() int {
	return rcvr.daysRented
}

func (rcvr Rental) Movie() Movie {
	return rcvr.movie
}

func (r Rental) Charge() float64 { //one param = one dependency
	return r.Movie().Charge(r.daysRented) //simple with func charge
}
