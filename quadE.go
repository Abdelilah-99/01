package piscine

import "github.com/01-edu/z01"

func QuadE(x, y int) {
	for j := 0; j < y; j++ {
		for i := 0; i < x; i++ {
			if j == 0 || j == y-1 {
				if i == 0 && j == 0 {
					z01.PrintRune('A')
				} else if i == x-1 && j == 0 {
					z01.PrintRune('C')
				} else if i == 0 && j == y-1 {
					z01.PrintRune('C')
				} else if i == x-1 && j == y-1 {
					z01.PrintRune('A')
				} else {
					z01.PrintRune('B')
				}
			} else {
				if i == 0 || i == x-1 {
					z01.PrintRune('B')
				} else {
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
	}
}
