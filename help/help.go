package help

import (
	"fmt"
)

//func update();
// func input();

func Graph(matr [3][3]int) {
	fmt.Printf("\nENTER A LINE(0-2) AND COLUMN(0-2)((USE ENTER :-) )) THEN MAKE UR CHOICE (X OR 0)!\n\n")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if matr[i][j] == 1 {
				fmt.Print(" X ")
			} else if matr[i][j] == 2 {
				fmt.Print(" 0 ")
			} else {
				fmt.Print("   ")
			}
			if j < 2 {
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

func Input() (line int, col int, choice string) {

	fmt.Scanln(&col)
	fmt.Scanln(&line)
	fmt.Scan(&choice)
	return
}

func Update(matr *[3][3]int, line int, col int, player int) bool {
	if line < 0 || line > 2 || col < 0 || col > 2 {
		fmt.Println("ERROR. ENTER A VALID COORDS")
		return false
	}
	if matr[col][line] != 0 {
		fmt.Printf("EROR// CELL IS ALREADY FULL!\n")
		return false
	}
	matr[line][col] = player
	return true
}
