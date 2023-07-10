package main

import (
	"fmt"
	"strings"
)

// При создании больших строк есть потенциальная угроза неэффективного использования памяти, так как при каждой итерации мы будем вынуждены пересоздавать строку снова и снова.
// Чтобы решить данную проблему, мы можем использоваться builder, буфер которого оптимизирует работу с памятью.

// Еще одна потенциальная проблема заключается в том, что делая срез в 100 элементов мы делаем срез в 100 байт. Для ASCII это не критично, но для Unicode проблема. Чтобы избежать
// ее, обрежем слайс руню

func someFunc() {
	v := createHugeString(1 << 10)
	justString := v[:100]
	str := []rune(v)
	justR := str[:100]
	fmt.Println(justString)
	fmt.Println(string(justR))

}

func main() {
	someFunc()

}

func createHugeString(length int) string {
	var builder strings.Builder
	for i := 0; i < length; i++ {
		builder.WriteString("Б")
	}
	return builder.String()
}
