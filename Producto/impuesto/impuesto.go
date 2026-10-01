package impuesto

func Impu(montoo float64, ivaa float64) float64 {

	iva := ivaa / 100
	monto := montoo * iva
	precfinal := montoo + monto
	return precfinal
}
