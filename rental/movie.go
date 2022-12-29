package rental

const (
	_ = iota //index runner start with 0 _ is to skip 0
	CHILDRENS
	NEW_RELEASE
	REGULAR
)

type Pricer interface {
	Charge(DaysRented int) float64
	PriceCode() int
}

type Children struct {
	priceCode int
}

func (c Children) Charge(daysRented int) float64 {
	result := 1.5
	if daysRented > 3 {
		result += float64(daysRented-3) * 1.5
	}
	return result
}

func (c Children) PriceCode() int {
	return c.priceCode
}

func CreateChildren() Children {
	return Children{
		priceCode: CHILDRENS,
	}
}

type Regular struct {
	priceCode int
}

func (r Regular) Charge(daysRented int) float64 {
	result := 2.0
	if daysRented > 2 {
		result += float64(daysRented-2) * 1.5
	}
	return result
}

func (r Regular) PriceCode() int {
	return r.priceCode
}

func CreateRegular() Regular {
	return Regular{
		priceCode: REGULAR,
	}
}

type NewRelease struct {
	priceCode int
}

func (n NewRelease) Charge(daysRented int) float64 {
	return float64(daysRented) * 3.0
}

func (n NewRelease) PriceCode() int {
	return n.priceCode
}

func CreateNewRelease() NewRelease {
	return NewRelease{
		priceCode: NEW_RELEASE,
	}
}

type Movie struct {
	title     string
	priceCode int
	Price   Pricer
}

func NewM(title string, charger Pricer) (rcvr Movie) {
	return Movie{
		title:     title,
		priceCode: charger.PriceCode(),
		Price:   charger,
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
