package utils

import (
	"github.com/shopspring/decimal"
)

type XDecimal struct {
	f decimal.Decimal
}

func Decimal(f float64) XDecimal {
	return XDecimal{f: decimal.NewFromFloat(f)}
}

func (x XDecimal) String() string {
	return x.f.String()
}

func (x XDecimal) Add(y float64) XDecimal {
	x.f = x.f.Add(decimal.NewFromFloat(y))
	return x
}

func (x XDecimal) Sub(y float64) XDecimal {
	x.f = x.f.Sub(decimal.NewFromFloat(y))
	return x
}

func (x XDecimal) Mul(y float64) XDecimal {
	x.f = x.f.Mul(decimal.NewFromFloat(y))
	return x
}

func (x XDecimal) Div(y float64) XDecimal {
	x.f = x.f.Div(decimal.NewFromFloat(y))
	return x
}

func (x XDecimal) Pow(y float64) XDecimal {
	x.f = x.f.Pow(decimal.NewFromFloat(y))
	return x
}

func (x XDecimal) Float() float64 {
	f, _ := x.f.Float64()
	return f
}

func (x XDecimal) FloatX() XFloat {
	return XFloat(x.Float())
}

func (x XDecimal) IsZero() bool {
	return x.f.IsZero()
}
