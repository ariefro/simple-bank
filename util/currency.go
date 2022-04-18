package util

const (
	USD = "USD"
	SGD = "SGD"
	EUR = "EUR"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, SGD:
		return true
	}
	return false
}