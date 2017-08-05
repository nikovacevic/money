// Package money represents money in a stable, integer-based way.
package money

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

// USD represents USD currecny
type USD int64

// NewUSD returns a new zero-value instance of USD
func NewUSD() *USD {
	m := USD(0)
	return &m
}

// ToUSD converts a float64 to USD, rounding to two decimal places
func ToUSD(f float64) USD {
	return USD((f * 100) + 0.5)
}

// Float64 converts USD to float64
func (m USD) Float64() float64 {
	x := float64(m)
	x = x / 100
	return x
}

// Multiply returns the USD product of a USD value with a float64
func (m USD) Multiply(f float64) USD {
	x := (float64(m) * f) + 0.5
	return USD(x)
}

// Scan attempts to scan a value into USD
func (m *USD) Scan(val interface{}) error {
	var v float64
	var err error
	switch val.(type) {
	case string:
		v, err = strconv.ParseFloat(val.(string), 64)
		if err != nil {
			return fmt.Errorf("Invalid string passed to Scan(). \"%v\" cannot be parsed to a USD value.", val)
		}
	case float64:
		v = val.(float64)
	case int:
		v = float64(val.(int))
	case []byte:
		s := string(val.([]byte))
		v, err = strconv.ParseFloat(s, 64)
		if err != nil {
			return fmt.Errorf("Invalid data passed to Scan(). \"%v\" cannot be parsed to a USD value.", s)
		}
	}
	*m = ToUSD(v)
	return nil
}

// String returns a formatted USD value
func (m USD) String() string {
	x := float64(m)
	x = x / 100
	return fmt.Sprintf("%.2f", x)
}

// Value converts USD to a DB-friendly value. Although USD is stored in cents,
// the Value will be a two-decimal numeric type; e.g. $18.00 is stored under
// the hood by USD as 1800, but Value() would return 18.00 for the DB.
func (m USD) Value() (driver.Value, error) {
	return m.Float64(), nil
}
