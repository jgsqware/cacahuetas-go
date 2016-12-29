package main

import (
	"fmt"
)

func main() {
	givers := map[string]string{
		"Julien Garcia Gonzalez": "garciagonzalez.julien@gmail.com",
		"Céline Liurno":          "celine.liurno@gmail.com",
		"Ludivine":               "ludivine@gmail.com",
		"Jeremie":                "jeremie@gmail.com",
		"Anthony Ennen":          "anthony.ennen@gmail.com",
		"Elisa Bono":             "elisa.bono@gmail.com",
	}

	receivers := copy(givers)

	fmt.Printf("givers: %v\n", givers)
	fmt.Printf("receivers: %v\n", receivers)
	fmt.Println("------")
	fmt.Printf("givers size: %v", len(givers))
}

func copy(originalMap map[string]string) (newMap map[string]string) {
	newMap = make(map[string]string)
	for k, v := range originalMap {
		newMap[k] = v
	}

	return
}
