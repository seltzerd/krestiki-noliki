package help

import (
	"fmt"
)

//func update();
// func input();

func Graph(matr [3][3]int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf(" %d ", matr[i][j])
			if j < 3-1 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if i < 3-1 {
			for j := 0; j < 3; j++ {
				fmt.Print("---")
				if j < 3-1 {
					fmt.Print("+")
				}
			}
			fmt.Println()
		}
	}
}

func Input(matr [3][3]int) {

}

func Update(matr [3][3]int) {

}
