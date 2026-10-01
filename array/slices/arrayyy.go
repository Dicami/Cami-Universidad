package main

import "fmt"

func main() {
	//arro()
	//slicee()
	//numerosa := []string{"manzana", "pera", "uva", "Teamosamuelito"}

	//for i, frutis := range numerosa {
	//fmt.Println("En el índice", i, "está la fruta:", frutis)
	//}
	//ejemp()
	//numeros := []string{"manzana", "pera", "uva"}
	//for _, fruta := range numeros {
	//fmt.Println("Fruta:", fruta)
	//}
	numeros := [3]int{1, 2, 3}
	for _, orden := range numeros {
		fmt.Println("El orden de los numeros es:", orden)
	}
}

func arro() {
	arra := [4]int{}
	for i := 0; i < 4; i++ {
		fmt.Println("Ingresesa los numeros ", i+1)
		fmt.Scan(&arra[i])
		fmt.Println(arra)
	}

}
func slicee() {
	slid := []int{}
	var cant int
	fmt.Println("Ingrese los valores que quiera digitar")
	fmt.Scan(&cant)
	for i := 0; i < cant; i++ {
		var ingre int
		fmt.Println("Ingrese la cantidad:")
		fmt.Scan(&ingre)
		slid = append(slid, ingre)
		fmt.Println(slid)

	}
}
func ejemp() {
	names := [4]string{"S4m_jn", "camiircs", "saramia_07", "minaooo"}
	for i, instas := range names {
		fmt.Println("Los indices son:", i, "y sus valores son:", instas)
	}

}
