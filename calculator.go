package mycalculator

import (
	"strconv"
	"strings"
	"os"
	"fmt"
	"bufio"
)

type Calc struct {}

func (Calc) Operate(input string, operator string) int {
	cleanInput := strings.Split(input, operator)
	number1 := parser(cleanInput[0])
	number2 := parser(cleanInput[1])

	switch operator {
		case "+":
			return number1 + number2

		case "-":
			return number1 - number2

		case "*":
			return number1 * number2

		case "/":
			return number1 / number2
		
		default:
			fmt.Println("Invalid input")
			return 0
	}
}

func parser(input string) int {
	number, _ := strconv.Atoi(input)

	return number
}

func ReadInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	
	return scanner.Text()
}

func GetOperator(operation string) string {
	var operator string

	if strings.Contains(operation, "+") {
		operator = "+"
	}
	if strings.Contains(operation, "-"){
		operator = "-"
	}
	if strings.Contains(operation, "*"){
		operator = "*"
	}
	if strings.Contains(operation, "/"){
		operator = "/"
	}
	
	return operator
}



