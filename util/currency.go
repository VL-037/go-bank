package util

const (
	IDR = "IDR"
	USD = "USD"
	EUR = "EUR"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case IDR, USD, EUR:
		return true
	}
	return false
}