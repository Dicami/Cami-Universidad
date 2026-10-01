package producto

func Produc(pro float64, des float64) float64 {
	montdes := pro * (des / 100)
	preciofinal := pro - montdes
	return preciofinal
}
