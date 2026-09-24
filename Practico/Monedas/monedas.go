package monedas

func ConvEuro(dolares float64) float64 {
	return dolares * 0.92
}

func ConvLibra(dolares float64) float64 {
	return dolares * 0.77
}

func ConvWon(dolares float64) float64 {
	return dolares * 1330.0
}

func ConvBTC(dolares float64) float64 {
	const precioBTC = 84805.0
	return dolares / precioBTC
}
