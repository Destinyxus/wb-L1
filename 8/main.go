package main

import "fmt"

func setBit(num int64, i uint, bitValue int) int64 {

	// 55 = 00110111

	// Используем побитовый сдвиг единицы влево на i-й битов, для того, чтобы получить маску для операции замены в исходном числе бита на единицу.
	//00000001 -  00001000 (если i == 3)
	bitmaskForAdd := int64(1) << i
	// Создаем маску для установления нуля на определенном месте исходного значения.
	//00000001 -  11011111 (если i == 5)
	bitmaskForClear := ^(int64(1) << i)

	if bitValue == 1 {
		num |= bitmaskForAdd
		// num = 00111111
	} else {
		num &= bitmaskForClear
	}

	return num
}

func main() {
	num := int64(55)

	num = setBit(num, 3, 1)
	fmt.Println(num)

	num = setBit(num, 5, 0)
	fmt.Println(num)

}
