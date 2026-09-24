package contador

func ContarVocales(texto string) (int, int, int, int, int) {
	var a, e, i, o, u int

	for _, letra := range texto {
		switch letra {
		case 'a', 'A':
			a++
		case 'e', 'E':
			e++
		case 'i', 'I':
			i++
		case 'o', 'O':
			o++
		case 'u', 'U':
			u++
		}
	}

	return a, e, i, o, u
}
