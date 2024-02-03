package main

func main() {
	/*
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
						casado := true

						fmt.Print("Hello, ")
						fmt.Print("world! \n")
						fmt.Print("!!\n")
						fmt.Print("My age is", age, "and my name is", name, "\n")

						fmt.Printf("My age is %v and my name is %v \n", age, name)                              //%v es el generico,toma el valor por defecto de la variable
						fmt.Printf("My age is %v and my name is %s \n", age, name)                              //%s es para interpolar un string
						fmt.Printf("My age is %d and my name is %v \n", age, name)                              //%d interpolar un entero en lugar de que salga como binario
						fmt.Printf("My age is %v , my name is %v and  my salary is %f \n", age, name, salary)   //f interpolar un float
						fmt.Printf("My age is %v , my name is %v and  my salary is %.2f \n", age, name, salary) //f interpolar un float, el .2 lo redondea
						fmt.Printf("age is type %T \n", age)                                                    //%T tipo de dato de la variable
						fmt.Printf("My age is %+v \n", age)                                                     //%+v incluye el nombre de la variable en caso de ser una estructura
						fmt.Printf("age is type %T \n", age)                                                    //%T tipo de dato de la variable
						fmt.Printf("is marry %t \n", casado)                                                    //%t interpolar booleano
						fmt.Printf(" age in binary code %b \n", age)                                            //%b representacion binaria
						fmt.Printf("my age in char is %c \n", age)                                              //%c caracter correspondiente a un numero
						fmt.Printf(" age in hex code %x \n", age)                                               //%b representacion hexadecimal
						fmt.Printf(" my name in double-quote %q \n", name)                                      //%q mostrar string entre comillas
						fmt.Printf(" age pointer is %p \n", &age)                                               //%p mostrar puntero
						fmt.Printf(" show  coquette age width:|%6d| \n", age)                                   //width:|%6d| mostrar la edad controlando el width en la vista
						s := fmt.Sprintf("creating string without saying I have %d years \n", age)              //save formatted strings
						fmt.Println(s)


					//Arrays y slice
					//var ages [3]int=[3]int{10,20,30}
					var ages = [3]int{10, 20, 30}
					names := [4]string{"irina", "ros", "carla", "fausto"}
					fmt.Println(ages, len(ages)) //tamanno logico del arreglo,no puede cambiar su tamanno
					fmt.Println(names, len(names))

					names[1] = "Luigi" //asignar valor a una posicion

					//slices (usa arrays por debajo)
					var scores = []int{10, 20, 30}
					scores[2] = 45
					scores = append(scores, 25) //append retorna un nuevo slice con el elemento guardado

					//silce range
					rangeOne := names[1:3] //obtener los valores entre la posicicion 1 y 3 sin incluir la 3
					rangeTwo := names[2:]  //obtener los valores de la posicion 2 al final
					rangeThre := names[:3] //obtener los valores del principio a la posicion 3 sin incluir la 3
					fmt.Println(rangeOne, rangeTwo, rangeThre)

				//Strings Package
				saludo := "Hola"
				fmt.Println(strings.Contains(saludo, "L"))       //buscar substring dentro de string
				hi := (strings.ReplaceAll(saludo, "Hola", "Hi")) //reemplazar substring,no modifica el original sino que crea uno nuevo
				fmt.Println(hi, saludo)
				fmt.Println(strings.ToUpper(saludo))        //convertir a mayuscula
				fmt.Println(strings.ToLower(saludo))        //convertir a minuscula
				fmt.Println(strings.Index(saludo, "do"))    //posicion donde comienza la cadena que buscamos
				saludoArray := (strings.Split(saludo, " ")) //split de toda la vida
				fmt.Println(saludoArray)


			//Sort
			ages := []int{20, 23, 14, 89}
			sort.Ints(ages) //ordenar asc,,,este metodo modifica el orden real en la slice
			fmt.Println(ages)
			index := sort.SearchInts(ages, 23)   //obtener posicion para un valor
			index2 := sort.SearchInts(ages, 100) //cuando no encuentra el valor devuelve len(ages)
			fmt.Println("Tamanno", len(ages))    //es como el size
			fmt.Println("Encontrado", index)
			fmt.Println(" No Encontrado", index2)

			names := []string{"a", "h", "b", "p"}
			sort.Strings(names) //ordenar lista de string
			fmt.Println(names)

			fmt.Println(sort.SearchStrings(names, "b")) //devuelve index de b dentro de names

		//For Loop
		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}
		names := []string{"a", "h", "b", "p"}
		for j := 0; j < len(names); j++ {
			fmt.Println(names[j])
		}

		for index, value := range names { //funciona como un forEach
			fmt.Printf("En el indice %v se encuentra el valor %v", index, value)

		}
		for _, value := range names { //en caso de que no necesite el indice,si no necesito el valor lo reemplazo con un _
			fmt.Printf("valores %v", value)

		}
	*/

}
