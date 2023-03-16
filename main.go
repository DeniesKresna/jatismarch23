package main

import (
	"fmt"
)

type Soal func()

func main() {
	Soals := []Soal{
		// Soal1,
		// Soal2,
		// Soal3,
		Soal4,
	}

	fmt.Printf("\n")
	for _, v := range Soals {
		v()
		fmt.Println("========================")
	}
}

func Soal1() {
	fmt.Printf("Soal1. Tukar variable tanpa variable ke tiga\n")
	a, b := 50, 63
	fmt.Printf("awal: a: %d, b: %d\n", a, b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("akhir: a: %d, b: %d\n", a, b)
}

func Soal2() {
	fmt.Printf("Soal2. Membalikkan string\n")
	input, output := "jatis", ""
	fmt.Printf("input: %s\n", input)
	for i := len(input) - 1; i >= 0; i-- {
		output += string(input[i])
	}
	fmt.Printf("output: %s\n", output)
}

func Soal3() {
	fmt.Printf("Soal3. Jumlah angka dan list character\n")
	input, output := "dani Maulana", ""
	fmt.Printf("input: %s\n", input)

	var (
		store     = make(map[rune]int)
		sortedKey = []rune{}
	)
	for _, v := range input {
		if v == 32 {
			continue
		}
		if v >= 65 && v <= 90 {
			if v == 77 {
			}
			v = v + 32
		}
		val, ok := store[v]
		if !ok {
			store[v] = 1
			sortedKey = append(sortedKey, v)
		} else {
			store[v] = val + 1
		}
	}
	for _, v := range sortedKey {
		addt := ""
		if store[v] > 1 {
			addt = fmt.Sprintf("%d%c", store[v], v)
		} else {
			addt = fmt.Sprintf("%c", v)
		}
		output += addt
	}

	fmt.Printf("output: %s\n", output)
}

func Soal4() {
	fmt.Printf("Soal4. Print bilangan, term khusus prima dan x9\n")
	fmt.Println("Output: ")

	isPrime := func(num int) bool {
		for i := 2; i < num; i++ {
			if 0 == num%i {
				return false
			}
		}
		return true
	}

	x9 := 1
	for i := 1; i <= 100; i++ {
		if 0 == i%9 {
			fmt.Println(i, "( Kelipatan 9 ke -", x9, ")")
			x9++
		} else if isPrime(i) {
			fmt.Println(i, "( Bilangan prima )")
		} else {
			fmt.Println(i)
		}
	}
}

func Soal5() {
	fmt.Printf("Soal2. Membalikkan string\n")
}
