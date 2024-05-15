package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("Введите вычисляемое выражение:")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка при вводе:", err)
		return
	}

	reg := regexp.MustCompile(`[+\-*/]`)
	parts := reg.Split(text, -1)
	for i, v := range parts {
		parts[i] = strings.TrimSpace(v)
	}

	fmt.Println(parts)
}
