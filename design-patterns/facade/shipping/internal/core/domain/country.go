package domain

type Country string

const (
	NG Country = "NG"
	US  Country = "US"
	UK  Country = "UK"
	UAE Country = "UAE"
	NIGERIA Country = "NG"
)

func (country Country) PrintCountry() string {
	switch country {
	case US:
		return string(US)
	case UK:
		return string(UK)
	case UAE:
		return string(UAE)
	case NG:
		return string(NG)
	}
	return "Unknown"
}
