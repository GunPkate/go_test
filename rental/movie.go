package rental

const (
	_ = iota //index runner start with 0 _ is to skip 0
	CHILDRENS
	NEW_RELEASE
	REGULAR
)

type Charger interface {
	Charge(DaysRented int) float64
}

type Children struct {
	// priceCode int
}

func (c Children) Charge(daysRented int) float64 {
	result := 1.5
	if daysRented > 3 {
		result += float64(daysRented-3) * 1.5
	}
	return result
}

type Movie struct {
	title     string
	priceCode int
	Charger   Charger
}

func NewM(title string, priceCode int, charger Charger) (rcvr Movie) {
	return Movie{
		title:     title,
		priceCode: priceCode,
		Charger:   charger,
	}
}

func NewMovie(title string, priceCode int) (rcvr Movie) {
	return Movie{
		title:     title,
		priceCode: priceCode,
	}

}
func (m Movie) PriceCode() int {
	return m.priceCode
}
func (m Movie) Title() string {
	return m.title
}
