package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
)

func main() {
	var numDecimal int = 42           // Десятичная система
	var numOctal int = 052            // Восьмеричная система
	var numHexadecimal int = 0x2A     // Шестнадцатиричная система
	var pi float64 = 3.14             // Тип float64
	var name string = "Golang"        // Тип string
	var isActive bool = true          // Тип bool
	var complexNum complex64 = 1 + 2i // Тип complex64

	Second(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	newStr := Third(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	newRunes := Fourth(newStr)
	fmt.Println(Fifths(newRunes))

}

func Second(dec, oct, hex int, fl float64, str string, bin bool, compl complex64) {
	fmt.Println("type of your number is :", GiveTypes(dec))
	fmt.Println("type of your number is :", GiveTypes(oct))
	fmt.Println("type of your number is :", GiveTypes(hex))
	fmt.Println("type of your number is :", GiveTypes(fl))
	fmt.Println("type of your number is :", GiveTypes(str))
	fmt.Println("type of your number is :", GiveTypes(bin))
	fmt.Println("type of your number is :", GiveTypes(compl))
}

func GiveTypes[T any](num T) reflect.Type {
	return reflect.TypeOf(num)
}

func Third(dec, oct, hex int, fl float64, str string, bin bool, compl complex64) string {
	return StringReturn(dec) + StringReturn(oct) + StringReturn(hex) + StringReturn(fl) + str + StringReturn(bin) + StringReturn(compl)
}

func StringReturn[T any](num T) string {
	return fmt.Sprintf("%v", num)
}

func Fourth(str string) []rune {
	return []rune(str)
}

func Fifths(runes []rune) string {
	hasher := sha256.New()
	var salt []rune
	salt = []rune("go-2024")
	var runesSalt []rune
	runesSalt = append(runesSalt, runes[0:len(runes)/2]...)
	runesSalt = append(runesSalt, salt...)
	runesSalt = append(runesSalt, runes[(len(runes)/2):]...)
	fmt.Println(string(runesSalt))
	hasher.Write([]byte(string(runesSalt)))
	return hex.EncodeToString(hasher.Sum(nil))
}
