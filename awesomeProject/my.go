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

		if want == "stop" {
			break
		}

		if want == "add" {
			fmt.Println("print num for search beebox:")
			var keyForMap int
			//	errors for key input
			_, err := fmt.Scan(&keyForMap)
			if err != nil {
				log.Fatal("the input num is not number:")
			}
			addInMap(keyForMap, bees)

		}

		if want == "watch" {
			fmt.Println(bees)
		}
		if want == "delete" {
			fmt.Println("print num beebox:")
			var keyForMap int
			//	errors for key input
			_, err := fmt.Scan(&keyForMap)
			if err != nil {
				log.Fatal("the input num is not number:")
			}
			delete(bees, keyForMap)
		}

	}
	fmt.Println(bees)
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
