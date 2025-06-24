package main

import (
	"blyat/help"
	"fmt"
)

func main() {
	var matr [3][3]int
	help.Graph(matr)
	for {
		winner := help.Win(matr)
		if winner == 1 {
			fmt.Printf("PLAYER 'X' WIN")
			break
		}
		if winner == 2 {
			fmt.Printf("PLAYER '0' WIN")
			break
		} else if help.Draw(matr) {
			fmt.Printf("DRAW")
			break
		}
		col, line, choice := help.Input()
		var player int
		if choice == "X" {
			player = 1
		} else if choice == "0" {
			player = 2
		} else {
			fmt.Println("ERROR")
			return
		}
		success := help.Update(&matr, line, col, player)
		if !success {
			fmt.Println("NOPE")
			return
		}
		help.Graph(matr)

	}

}
