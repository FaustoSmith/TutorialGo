package main

import "fmt"

func main() {

	//strings
	var nameone = "fausto"
	var nameTwo string = "carla"
	var nameThree string

	fmt.Println("Los nombres son", nameone, nameTwo, nameThree)

	nameone = "Pearl"
	nameThree = "browser"
	fmt.Println(nameone, nameThree)

	nameFour := "yosi"
	fmt.Println(nameFour)

	//ints
	var ageOne int = 8
	var agetwo = 30
	agethree := 80

	fmt.Println(ageOne, agetwo, agethree)

	//bits & memory
	//var numOne int8 = 25
	//var numTow int16 = -128
	//var numThree int32 = 255
	//var numFour uint16 = 90 //Posistve number

	//floats
	var scoreOne float32 = 45.5
	var scoreTwo float64 = 65.6
	scoreThree := 1.5

	fmt.Println(scoreOne, scoreTwo, scoreThree)

	//cons
	const secondsInMinute = 60
	const minutesInHour = 60
	fmt.Print(secondsInMinute, minutesInHour)

	//Printf
	age := 30
	name := "Fausto"
	salary := 1200.5378

	fmt.Print("Hello, ")
	fmt.Print("world! \n")
	fmt.Print("!!\n")
	fmt.Print("My age is", age, "and my name is", name, "\n")

	fmt.Printf("My age is %v and my name is %v", age, name)                              //%v es el generico,toma el valor por defecto de la variable
	fmt.Printf("My age is %v and my name is %s", age, name)                              //%s es para interpolar un string
	fmt.Printf("My age is %d and my name is %v", age, name)                              //%d interpolar un entero en lugar de que salga como binario
	fmt.Printf("My age is %v , my name is %v and  my salary is %f", age, name, salary)   //f interpolar un float
	fmt.Printf("My age is %v , my name is %v and  my salary is %.2f", age, name, salary) //f interpolar un float, el .2 lo redondea

}
