package mycalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readOperation() {
	// Declarar Scannner para leer desde el Teclado
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Escribe la operacion que desees separada con espacios ej: 2 + 2")

	// Leer entrada del teclado
	scanner.Scan()

	// Fragmentar la entrada por espacios [" "]
	operation := strings.Split(scanner.Text(), " ")

	// Obtener los numeros(convertirdos a enteros) y operador
	num1 := parseStrToInt(operation[0])
	op := operation[1]
	num2 := parseStrToInt(operation[2])

	// Instanciar nuestro struct [Calculator]
	cal := Calculator{}

	// Enviar los elementos de la operacion para ser resuelta
	cal.ResultOperation(op, num1, num2)
}

// Funcion privada que conviente un string a int
func parseStrToInt(num string) int {

	//Conversion del string con manejo de errores
	num1, err := strconv.Atoi(num)

	//En caso de error retornar 0
	if err != nil {
		return 0
	}

	//Regresar el numero convertido
	return num1
}

// Creamos nuestro Struct [Calculator]
type Calculator struct{}

// Funcion Receiver publica que recibe un operador y dos numeros y resuleve la operacion
func (Calculator) ResultOperation(op string, num1 int, num2 int) {

	// Condicionamos nuestro operador
	switch op {
	case "+":
		// Hacer suma
		fmt.Println("El resultado de la suma es:", (num1 + num2))

	case "-":
		// Hacer resta
		fmt.Println("El resultado de la resta es:", (num1 - num2))
		break

	case "*":
		// Hacer multiplicacion
		fmt.Println("El resultado de la multiplicacion es:", (num1 * num2))
		break

	case "/":
		// Hacer division
		fmt.Println("El resultado de la division es:", (num1 / num2))
		break

	default:
		// Indicar que no es soportada la operacion
		fmt.Println("El operador:[", op, "]no esta soportado")
		break
	}
}
