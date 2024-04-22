package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func IntegerToRoman(num int) string {
	max := 101
	if num > max {
		return strconv.Itoa(num)
	}
	conv := []struct {
		value int
		digit string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	var roman string
	for _, conv := range conv {
		for num >= conv.value {
			roman += conv.digit
			num -= conv.value
		}
	}
	return roman
}

var RomArab = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
}
var Operac = map[string]rune{
	"+": '+',
	"-": '-',
	"/": '/',
	"*": '*',
}
var oper string

func romanToInt(s string) int {
	var sum, big int
	for i := len(s) - 1; i >= 0; i-- {
		simbol := s[i]
		num := RomArab[rune(simbol)]
		if num >= big {
			big = num
			sum += num
			continue
		}
		sum -= num
	}
	return sum
}

var u string

const (
	chit = "Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."
	sos  = "Выдача паники, так как строка не является математической операцией."
	her  = "Выдача паники, так как в римской системе нет отрицательных чисел."
	not  = "Выдача паники, римлянам не нужен 0."
	hren = "Выдача паники, так как используются одновременно разные системы счисления."
	glob = "Используй  цифры от 1 до 10 включительно"
)

func main() {
	var y, z, q int
	reader := bufio.NewReader(os.Stdin)
	for {
		oper = u
		fmt.Println("input :")
		input, _ := reader.ReadString('\n')
		s := strings.ReplaceAll(input, " ", "")
		for znak := range Operac {
			for _, val := range s {
				if znak == string(val) {
					oper += znak
				}
			}
		}
		switch {
		case len(oper) > 1:
			panic(chit)
		case len(oper) < 1:
			panic(sos)
		}
		var a, b, w string
		for i := 0; i < len(s); i++ {
			simbol := s[i]
			if string(simbol) == oper {
				//b = s[i+1:] // БАГ
				break
			}
			a += string(simbol)
		}
		for i := len(s) - 1; i > 0; i-- {
			simbol := s[i]
			if string(simbol) == oper {
				break
			}
			w += string(simbol)
		}
		for i := len(w) - 1; i > 0; i-- {
			simbol := w[i]
			b += string(simbol)
		}
		z, q = romanToInt(a), romanToInt(b)
		if 10 < z || q > 10 {
			panic(glob)
		}
		if z == 0 && q > 0 || z > 0 && q == 0 {
			panic(hren)
		}
		if 0 < z && q > 0 {
			switch oper {
			case "+":
				y = z + q
			case "-":
				y = z - q
			case "*":
				y = z * q
			case "/":
				y = z / q
			}
			fmt.Println("output :")
			if y == 0 {
				panic(not)
			}
			if y <= 0 {
				panic(her)
			}
			fmt.Println(IntegerToRoman(y))
			continue
		}
		if n, err := strconv.Atoi(a); err == nil {
			z = n
			if z > 10 {
				panic(glob)
			}
		}
		if m, err := strconv.Atoi(b); err == nil {
			q = m
			if q > 10 {
				panic(glob)
			}
		}

		if z > 0 && q == 0 || q <= 0 && z <= 0 {
			panic(sos)
		}

		fmt.Println("output :")
		switch oper {
		case "+":
			y = z + q
		case "-":
			y = z - q
		case "*":
			y = z * q
		case "/":
			y = z / q
		}
		fmt.Println(y)
	}
}
