// temperatureConversion
// maritzavania

package main

import "fmt"

func main() {
	var celcius float64
	var choice int

	fmt.Println("Masukkan suhu dalam Celcius : ")
	fmt.Scanln(&celcius)

	fmt.Println("Pilih konversi satuan.")
	fmt.Println("1. Fahrenheit")
	fmt.Println("2. Kelvin")
	fmt.Println("3. Reamur")
	fmt.Println("Masukkan pilihan (1/2/3) : ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		var fahrenheit float64
		fahrenheit = (celcius * 9 / 5) + 32

		fmt.Println("Suhu dalam Fahrenheit adalah ", fahrenheit)
	case 2:
		var kelvin float64
		kelvin = celcius + 273

		fmt.Println("Suhu dalam Kelvin adalah ", kelvin)
	case 3:
		var reamur float64
		reamur = celcius * 4 / 5

		fmt.Println("Suhu dalam reamur adalah ", reamur)
	default:
		fmt.Println("Nomor tidak valid. Harap masukkan pilihan (1/2/3).")

	}

}
