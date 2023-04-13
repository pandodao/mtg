package mtgpack

import "github.com/shopspring/decimal"

type Number struct {
	decimal.Decimal
}

func (n Number) EncodeMtg(e *Encoder) error {
	return e.EncodeDecimal(n.Decimal)
}
