package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitlist = []string{"apple", "mango", "grapes"}
	fmt.Printf("%T", fruitlist)
	fruitlist = append(fruitlist, "orange", "banana")
	fmt.Println(fruitlist)
	fruitlist = append(fruitlist[:2])
	fmt.Println(fruitlist)

	highscores := make([]int, 4)

	highscores[0] = 234
	highscores[1] = 945
	highscores[2] = 434
	highscores[3] = 534
	//highscores[4] = 123 (error);

	//in append memory is allocated again
	highscores = append(highscores, 555, 666, 321)
	fmt.Println(highscores)
	sort.Ints(highscores)
	fmt.Println(highscores)

	//remove element from slice through index

	var ind int = 2
	highscores = append(highscores[:ind], highscores[ind+1:]...)
	fmt.Println(highscores)
}
