package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func makeOperation(operation string, firstInteger int, secondInteger int) (int, string) {
	switch operation {
	case "+":
		return firstInteger + secondInteger, ""
	case "-":
		return firstInteger - secondInteger, ""
	case "*":
		return firstInteger * secondInteger, ""
	case "/":
		return firstInteger / secondInteger, ""
	case "%":
		return firstInteger % secondInteger, ""
	default:
		return 0, "Esa operación no concuerda con ninguna de las anteriores"
	}
}

func getOperation(scanner *bufio.Scanner) string {
	scanner.Scan()
	operation := scanner.Text()
	return operation
}

func getValue(scanner *bufio.Scanner) (int, error) {
	scanner.Scan()
	scannedText := scanner.Text()
	value, err := strconv.Atoi(scannedText)
	return value, err
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Bienvenido a mi calculadora en Golang")
	finished := false

	for !finished {
		fmt.Println("Ingresa la operación que necesitas hacer:")
		fmt.Println("Suma: +, Resta: -, Multiplicación: *, División: / o %")
		operation := getOperation(scanner)
		fmt.Printf("Agrega el primer valor para la operación seleccionada %s \n", operation)
		firstValue, err := getValue(scanner)
		if err != nil {
			fmt.Print("Lo que ingresaste no es un número, volviendo al generar la calculdora...")
			continue
		}
		fmt.Printf("Agrega el segundo valor para la operación seleccionada %s \n", operation)
		secondValue, err2 := getValue(scanner)
		if err2 != nil {
			fmt.Println("Lo que ingresaste no es un número, volviendo al generar la calculdora...")
			continue
		}
		result, errorString := makeOperation(operation, firstValue, secondValue)
		fmt.Println("Resultado -> ", result)
		fmt.Println("¿Deseas continuar haciendo operaiones?")
		scanner.Scan()
		response := scanner.Text()
		if response == "no" {
			finished = true
			fmt.Println("Resultado final -> ", result, errorString)
		}
		continue
	}
	fmt.Println("La calculadora se está cerrando...")
}
