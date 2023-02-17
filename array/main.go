package main

import (
	"fmt"
)

func main() {
	var fruitList [4]string

	fruitList[0] = "apple"
	fruitList[1] = "tomato"
	fruitList[2] = "peach"
	fmt.Println("welcome to array in golangs")
	fmt.Println("fruit list is :", fruitList)
	fmt.Println("fruit list is :", len(fruitList))
	var vegLists = [3]string{"potato", "carrot", "mushroom"}
	fmt.Println("veggie list is: ", len(vegLists))

}
