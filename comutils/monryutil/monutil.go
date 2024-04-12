package monryutil

import "github.com/shopspring/decimal"

// FenToYuan 分转元
// fen: 分
//
// return: 元
func FenToYuan(fen int64) float64 {
	yuan := decimal.NewFromInt(fen).Div(decimal.NewFromInt(100)).Round(2)
	f, _ := yuan.Float64()
	return f
}

// YuanToFen 元转分
// yuan: 元
//
// return: 分
func YuanToFen(yuan float64) int64 {
	d := decimal.New(1, 2)
	fen := decimal.NewFromFloat(yuan).Mul(d).IntPart()
	return fen
}
