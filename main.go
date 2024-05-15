package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func partsOperations(parts, operator []string) string {
	var answer string
	switch operator[0] {
	default:
		fmt.Println("Что-то пошло не так...")
	case "+":
		answer = parts[0] + parts[1]
	case "-":
		answer = strings.Replace(parts[0], parts[1], "", 1)
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

	reg := regexp.MustCompile(`\"[+\-*/\s]+\"`)
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
		parts[i] = strings.Trim(v, `"`)
	}
	for i, v := range parts {
		parts[i] = strings.TrimSpace(v)
	}

	fmt.Println("дебаг. parts и operator", parts, operator)
	fmt.Println(partsOperations(parts, operator))

}
