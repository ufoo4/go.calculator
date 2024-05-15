package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func divideString(input string, n int) string {
	parts := make([]string, 0, n)
	length := len(input)
	partSize := length / n

	for i := 0; i < length; i += partSize {
		end := i + partSize
		if end > length {
			end = length
		}
		parts = append(parts, input[i:end])
	}
	return parts[0]
}

func partsOperations(parts, operator []string) string {
	var answer string
	switch operator[0] {
	default:
		fmt.Println("Что-то пошло не так...")
	case "+":
		answer = parts[0] + parts[1]
	case "-":
		answer = strings.Replace(parts[0], parts[1], "", 1)
	case "*":
		operand2, _ := strconv.Atoi(parts[1])
		answer = strings.Repeat(parts[0], operand2)
	case "/":
		operand2, _ := strconv.Atoi(parts[1])
		answer = divideString(parts[0], operand2)
	}
	return answer
}

func main() {
	fmt.Println("Введите вычисляемое выражение:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка при вводе:", err)
		return
	}

	reg := regexp.MustCompile(`\"[+\-*/\s]+`)
	parts := reg.Split(input, -1)
	operator := reg.FindAllString(input, -1)
	for i, v := range operator {
		operator[i] = strings.Trim(v, `"`)
	}
	for i, v := range operator {
		operator[i] = strings.TrimSpace(v)
	}

	if len(operator[0]) > 1 {
		fmt.Printf("Ошибка при вводе оператора %s в выражении\n", operator)
		return
	}
	for i, v := range parts {
		parts[i] = strings.TrimSpace(v)
	}
	for i, v := range parts {
		parts[i] = strings.Trim(v, `"`)
	}

	fmt.Printf("\"%s\"\n", partsOperations(parts, operator))

}
