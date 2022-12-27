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
func (rcvr Rental) GetDaysRented() int {
	return rcvr.daysRented
}
func (rcvr Rental) GetMovie() Movie {
	return rcvr.movie
}
