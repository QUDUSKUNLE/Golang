package facade

import "fmt"


type SecurityCode struct {
	code int
}

func NewSecurityCode(code int) *SecurityCode {
	return &SecurityCode{
		code: code,
	}
}

func (security *SecurityCode) CheckSecurityCode(code int) error {
	if security.code != code {
		return fmt.Errorf("security code is incorrect")
	}
	fmt.Println("Security code is verified")
	return nil
}
