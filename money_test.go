package money_test

import (
	"database/sql/driver"
	"fmt"
	"testing"

	"github.com/nikovacevic/money"
)

var scanTests = []struct {
	val interface{}
	exp money.USD
	err error
}{
	{0, money.USD(0), nil},
	{10, money.USD(1000), nil},
	{1.23, money.USD(123), nil},
	{1.345, money.USD(135), nil},
	{"10", money.USD(1000), nil},
	{"1.345", money.USD(135), nil},
	{[]uint8{49, 56, 46, 48, 48}, money.USD(1800), nil},
	{[]uint8{49, 50, 46, 115, 100}, money.USD(0), fmt.Errorf("Invalid data passed to Scan(). \"%v\" cannot be parsed to a USD value.", "12.sd")},
	{"1O", money.USD(0), fmt.Errorf("Invalid string passed to Scan(). \"%v\" cannot be parsed to a USD value.", "1O")},
}

func TestScan(t *testing.T) {
	for _, test := range scanTests {
		m := money.NewUSD()
		err := m.Scan(test.val)
		if fmt.Sprintf("%s", err) != fmt.Sprintf("%s", test.err) {
			t.Errorf("Scan() returned error \"%v\". Should return error \"%v\".", err, test.err)
		}
		if *m != test.exp {
			t.Errorf("Scan() returned %v. Should return %v.", m, test.exp)
		}
	}
}

var valueTests = []struct {
	val money.USD
	exp driver.Value
	err error
}{
	{money.USD(0), 0.00, nil},
	{money.USD(1850), 18.50, nil},
}

func TestValue(t *testing.T) {
	for _, test := range valueTests {
		drv, err := test.val.Value()
		if fmt.Sprintf("%s", err) != fmt.Sprintf("%s", test.err) {
			t.Errorf("Scan() returned error \"%v\". Should return error \"%v\".", err, test.err)
		}
		if drv.(float64) != test.exp.(float64) {
			t.Errorf("Scan() returned %v. Should return %v.", drv, test.exp)
		}
	}
}

var floatTests = []struct {
	m   money.USD
	exp float64
}{
	{money.USD(0), 0.0},
	{money.USD(1), 0.01},
	{money.USD(12345), 123.45},
}

func TestFloat64(t *testing.T) {
	for _, test := range floatTests {
		f := test.m.Float64()
		if f != test.exp {
			t.Errorf("(%v).Float64() returned %v. Should return %v.", test.m, f, test.exp)
		}
	}
}

var multiplyTests = []struct {
	m   money.USD
	x   float64
	exp money.USD
}{
	{money.USD(1), 1.0, money.USD(1)},
	{money.USD(1234), 1.0, money.USD(1234)},
	{money.USD(1000), 0.05, money.USD(50)},
	{money.USD(1000), 0.0555, money.USD(56)},
}

func TestMultiply(t *testing.T) {
	for _, test := range multiplyTests {
		p := test.m.Multiply(test.x)
		if p != test.exp {
			t.Errorf("(%v).Multiply(%v) returned %v. Should return %v.", test.m, test.x, p, test.exp)
		}
	}
}

var stringTests = []struct {
	m   money.USD
	exp string
}{
	{money.USD(1), "$0.01"},
	{money.USD(1234), "$12.34"},
}

func TestString(t *testing.T) {
	for _, test := range stringTests {
		s := test.m.String()
		if s != test.exp {
			t.Errorf("(%v).String() returned %v. Should return %v.", test.m, s, test.exp)
		}
	}
}
