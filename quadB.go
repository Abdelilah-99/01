package piscine

import "github.com/01-edu/z01"

func QuadB(x, y int) {
	for j := 0; j < y; j++ {
		for i := 0; i < x; i++ {
			if j == 0 || j == y-1 {
				if i == 0 && j == 0 {
					z01.PrintRune('/')
				} else if i == x-1 && j == 0 {
					z01.PrintRune('\\')
				} else if i == 0 && j == y-1 {
					z01.PrintRune('\\')
				} else if i == x-1 && j == y-1 {
					z01.PrintRune('/')
				} else {
					z01.PrintRune('*')
				}
			} else {
				if i == 0 || i == x-1 {
					z01.PrintRune('*')
				} else {
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
	}
}
