package interfaces

import "fmt"

type PaymentMethod interface {
	Process(amount float64) bool
	Provider() string
}

type CardPayment struct {
	CardNumber string
	Limit      float64
}

func (c *CardPayment) Process(amount float64) bool {
	if amount <= c.Limit {
		c.Limit -= amount
		return true
	}
	return false
}

func (c CardPayment) Provider() string {
	return "CARD"
}

type UPIPayment struct {
	VPA string
}

func (u UPIPayment) Process(amount float64) bool {
	return true
}

func (u UPIPayment) Provider() string {
	return "UPI"
}

type CryptoPayment struct {
	Wallet  string
	Balance float64
}

func (c *CryptoPayment) Process(amount float64) bool {
	if amount <= c.Balance {
		c.Balance -= amount
		return true
	}
	return false
}

func (c CryptoPayment) Provider() string {
	return "CRYPTO"
}

func Checkout(p PaymentMethod, amount float64) string {
	if ok := p.Process(amount); ok {
		return fmt.Sprintf("Payment successful via %s", p.Provider())
	}
	return fmt.Sprintf("Payment failed via %s", p.Provider())
}

func DetectCrypto(p PaymentMethod) bool {
	if _, ok := p.(*CryptoPayment); ok {
		return true
	}
	return false
}
