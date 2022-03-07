package main

import "fmt"

func main() {
	var nilai1 int16
	var nilai2 int16
	var ratarata int16

	fmt.Println("Masukan Nilai 1 :")
	fmt.Scanln(&nilai1)

	fmt.Println("Masukan Nilai 2 :")
	fmt.Scanln(&nilai2)

	fmt.Println("Masukan rata - rata :")
	fmt.Scanln(&ratarata)

	if (nilai1+nilai2)/2 >= ratarata {
		fmt.Println("Nilai Tuntas")
	} else if (nilai1+nilai2)/2 < ratarata {
		fmt.Println("Remidi")
	}

}
