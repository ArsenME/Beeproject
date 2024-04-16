package main

import (
	"fmt"
	"log"
)

func main() {

	bees := make(map[int]string)
	for {
		var want string
		fmt.Println("how do you want?:")
		//errors
		_, err := fmt.Scanf("%s", &want)
		if err != nil {
			log.Fatal("error reading input")
		}
		switch want {
		case "add":
			fmt.Println("print num for search beebox:")
			var keyForMap int
			//	errors for key input
			_, err := fmt.Scan(&keyForMap)
			if err != nil {
				log.Fatal("the input num is not number:")
			}
			addInMap(keyForMap, bees)
		case "delete":
			deleteFromMap(bees)

		case "watch":
			fmt.Println(bees)

		case "stop":
			break
		default:
			break
		}
		fmt.Println(bees)
	}

}

func addInMap(x int, y map[int]string) {
	fmt.Println("input subjects: ")
	s := ""
	_, err := fmt.Scanln(&s)
	if err != nil {
		log.Fatal("Error reading input:", err)
	}
	y[x] = s

}

func deleteFromMap(y map[int]string) {
	fmt.Println("print num beebox:")
	var keyForMap int
	//	errors for key input
	_, err := fmt.Scan(&keyForMap)
	if err != nil {
		log.Fatal("the input num is not number:")
	}
	delete(y, keyForMap)
}
