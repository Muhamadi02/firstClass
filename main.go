package main

import (
	"fmt"
	"strings"
)

func main() {

	myString := strings.Split("Я сейчас пойду домой", " ")

	fmt.Println(myString[0])

	newString := strings.Replace("Я узучаю язык Golang", "язык", "language", -1)
	fmt.Println(newString)

	if strings.Contains("Я иду домой", "иду") {
		fmt.Println("contains")
	} else {
		fmt.Println("doesn't contain")
	}

	/*var xyz float64
	var someFloat string
	xyz = 1.2
	someFloat = "2.1"

	newFloat, err := strconv.ParseFloat(someFloat, 64)
	if err != nil {
		fmt.Println("parse err: ", err.Error())
		return
	}

	fmt.Println(xyz + newFloat)

	fmt.Printf("Sum: %v", xyz+newFloat)*/

	/*var xyz int
	var someFloat float64
	xyz = 1
	someFloat = 2.1

	fmt.Println(float64(xyz) + someFloat)
	fmt.Println((float64(xyz) + someFloat) / 2)*/

	/*var xyz int

	fmt.Print("Вводите число: ")
	_, err := fmt.Scan(&xyz)
	if err != nil {
		fmt.Println("Ошибка при чтении из терминала: ", err.Error())
		return
	}

	if _, err := fmt.Scan(&xyz); err != nil {
		fmt.Println("Ошибка при чтении из терминала: ", err.Error())
		return
	}*/
	//fmt.Println(err.Error())
}

type Human struct {
	Head  interface{}
	Hands int8
	Feet  int8
	Torso []string
}
